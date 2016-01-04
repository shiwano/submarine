using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using ModestTree;
using ModestTree.Util;

#if !ZEN_NOT_UNITY3D
using UnityEngine;
#endif

namespace Zenject
{
    // Responsibilities:
    // - Expose methods to configure object graph via Bind() methods
    // - Build object graphs via Resolve() method
    public class DiContainer : IInstantiator, IResolver, IBinder
    {
        readonly Dictionary<BindingId, List<ProviderBase>> _providers = new Dictionary<BindingId, List<ProviderBase>>();
        readonly SingletonProviderMap _singletonMap;
        readonly HashSet<Type> _installedInstallers = new HashSet<Type>();
        readonly Stack<Type> _installsInProgress = new Stack<Type>();
        readonly DiContainer _parentContainer;
        readonly Stack<LookupId> _resolvesInProgress = new Stack<LookupId>();

#if !ZEN_NOT_UNITY3D
        readonly Transform _rootTransform;
#endif

        bool _isValidating;

        public DiContainer()
        {
            _singletonMap = new SingletonProviderMap(this);

            this.Bind<DiContainer>().ToInstance(this);
            this.Bind<IInstantiator>().ToInstance(this);
            this.Bind<SingletonProviderMap>().ToInstance(_singletonMap);

#if !ZEN_NOT_UNITY3D
            this.Bind<PrefabSingletonProviderMap>().ToSingle<PrefabSingletonProviderMap>();
#endif
            this.Bind<SingletonInstanceHelper>().ToSingle<SingletonInstanceHelper>();
        }

        public DiContainer(DiContainer parentContainer)
            : this()
        {
            _parentContainer = parentContainer;
        }

#if !ZEN_NOT_UNITY3D
        public DiContainer(Transform rootTransform, DiContainer parentContainer)
            : this()
        {
            _parentContainer = parentContainer;
            _rootTransform = rootTransform;
        }

        public DiContainer(Transform rootTransform)
            : this()
        {
            _rootTransform = rootTransform;
        }
#endif

        public SingletonProviderMap SingletonProviderMap
        {
            get
            {
                return _singletonMap;
            }
        }

        public DiContainer ParentContainer
        {
            get
            {
                return _parentContainer;
            }
        }

        public bool ChecksForCircularDependencies
        {
            get
            {
#if ZEN_MULTITHREADING
                // When multithreading is supported we can't use a static field to track the lookup
                // TODO: We could look at the inject context though
                return false;
#else
                return true;
#endif
            }
        }

        public IEnumerable<Type> InstalledInstallers
        {
            get
            {
                return _installedInstallers;
            }
        }

#if !ZEN_NOT_UNITY3D
        public Transform RootTransform
        {
            get
            {
                return _rootTransform;
            }
        }
#endif

        // True if this container was created for the purposes of validation
        // Useful to avoid instantiating things that we shouldn't during this step
        public bool IsValidating
        {
            get
            {
                return _isValidating;
            }
            set
            {
                _isValidating = value;
            }
        }

        public IEnumerable<BindingId> AllContracts
        {
            get
            {
                return _providers.Keys;
            }
        }

        // Note that this list is not exhaustive or even accurate so use with caution
        public IEnumerable<Type> AllConcreteTypes
        {
            get
            {
                return (from x in _providers from p in x.Value select p.GetInstanceType()).Where(x => x != null && !x.IsInterface && !x.IsAbstract).Distinct();
            }
        }

        public DiContainer CreateSubContainer()
        {
#if ZEN_NOT_UNITY3D
            return new DiContainer(this);
#else
            return new DiContainer(_rootTransform, this);
#endif
        }

        public void RegisterProvider(
            ProviderBase provider, BindingId bindingId)
        {
            if (_providers.ContainsKey(bindingId))
            {
                // Prevent duplicate singleton bindings:
                if (_providers[bindingId].Find(item => ReferenceEquals(item, provider)) != null)
                {
                    throw new ZenjectBindException(
                        "Found duplicate singleton binding for contract '{0}' and id '{1}'".Fmt(bindingId.Type, bindingId.Identifier));
                }

                _providers[bindingId].Add(provider);
            }
            else
            {
                _providers.Add(bindingId, new List<ProviderBase> {provider});
            }
        }

        public int UnregisterProvider(ProviderBase provider)
        {
            int numRemoved = 0;

            foreach (var keyValue in _providers)
            {
                numRemoved += keyValue.Value.RemoveAll(x => x == provider);
            }

            Assert.That(numRemoved > 0, "Tried to unregister provider that was not registered");

            // Remove any empty contracts
            foreach (var bindingId in _providers.Where(x => x.Value.IsEmpty()).Select(x => x.Key).ToList())
            {
                _providers.Remove(bindingId);
            }

            provider.Dispose();

            return numRemoved;
        }

        public IEnumerable<Type> GetDependencyContracts<TContract>()
        {
            return GetDependencyContracts(typeof(TContract));
        }

        public IEnumerable<ZenjectResolveException> ValidateResolve<TContract>()
        {
            return ValidateResolve<TContract>((string)null);
        }

        public IEnumerable<ZenjectResolveException> ValidateResolve<TContract>(string identifier)
        {
            return ValidateResolve(new InjectContext(this, typeof(TContract), identifier));
        }

        public IEnumerable<ZenjectResolveException> ValidateValidatables(params Type[] ignoreTypes)
        {
            // Use ToList() in case it changes somehow during iteration
            foreach (var pair in _providers.ToList())
            {
                var bindingId = pair.Key;

                if (ignoreTypes.Where(i => bindingId.Type.DerivesFromOrEqual(i)).Any())
                {
                    continue;
                }

                // Validate all IValidatableFactory's
                List<ProviderBase> validatableFactoryProviders;

                var providers = pair.Value;

                if (bindingId.Type.DerivesFrom<IValidatableFactory>())
                {
                    validatableFactoryProviders = providers;
                }
                else
                {
                    validatableFactoryProviders = providers.Where(x => x.GetInstanceType().DerivesFrom<IValidatableFactory>()).ToList();
                }

                var injectCtx = new InjectContext(this, bindingId.Type, bindingId.Identifier);

                foreach (var provider in validatableFactoryProviders)
                {
                    var factory = (IValidatableFactory)provider.GetInstance(injectCtx);

                    var type = factory.ConstructedType;
                    var providedArgs = factory.ProvidedTypes;

                    foreach (var error in ValidateObjectGraph(type, injectCtx, null, providedArgs))
                    {
                        yield return error;
                    }
                }

                // Validate all IValidatable's
                List<ProviderBase> validatableProviders;

                if (bindingId.Type.DerivesFrom<IValidatable>())
                {
                    validatableProviders = providers;
                }
                else
                {
                    validatableProviders = providers.Where(x => x.GetInstanceType().DerivesFrom<IValidatable>()).ToList();
                }

                Assert.That(validatableFactoryProviders.Intersect(validatableProviders).IsEmpty(),
                    "Found provider implementing both IValidatable and IValidatableFactory.  This is not allowed.");

                foreach (var provider in validatableProviders)
                {
                    var factory = (IValidatable)provider.GetInstance(injectCtx);

                    foreach (var error in factory.Validate())
                    {
                        yield return error;
                    }
                }
            }
        }

        // Wrap IEnumerable<> to avoid LINQ mistakes
        internal List<ProviderBase> GetAllProviderMatches(InjectContext context)
        {
            return GetProviderMatchesInternal(context).Select(x => x.Provider).ToList();
        }

        // Be careful with this method since it is a coroutine
        IEnumerable<ProviderPair> GetProviderMatchesInternal(InjectContext context)
        {
            return GetProvidersForContract(context.BindingId, context.LocalOnly).Where(x => x.Provider.Matches(context));
        }

        IEnumerable<ProviderPair> GetProvidersForContract(BindingId bindingId, bool localOnly)
        {
            var localPairs = GetLocalProviders(bindingId).Select(x => new ProviderPair(x, this));

            if (localOnly || _parentContainer == null)
            {
                return localPairs;
            }

            return localPairs.Concat(
                _parentContainer.GetProvidersForContract(bindingId, false));
        }

        List<ProviderBase> GetLocalProviders(BindingId bindingId)
        {
            List<ProviderBase> localProviders;

            if (_providers.TryGetValue(bindingId, out localProviders))
            {
                return localProviders;
            }

            // If we are asking for a List<int>, we should also match for any localProviders that are bound to the open generic type List<>
            // Currently it only matches one and not the other - not totally sure if this is better than returning both
            if (bindingId.Type.IsGenericType && _providers.TryGetValue(new BindingId(bindingId.Type.GetGenericTypeDefinition(), bindingId.Identifier), out localProviders))
            {
                return localProviders;
            }

            return new List<ProviderBase>();
        }

        public IList ResolveAll(InjectContext context)
        {
            // Note that different types can map to the same provider (eg. a base type to a concrete class and a concrete class to itself)

            var matches = GetProviderMatchesInternal(context).ToList();

            if (matches.Any())
            {
                return ReflectionUtil.CreateGenericList(
                    context.MemberType, matches.Select(x => SafeGetInstance(x.Provider, context)).ToArray());
            }

            if (!context.Optional)
            {
                throw new ZenjectResolveException(
                    "Could not find required dependency with type '" + context.MemberType.Name() + "' \nObject graph:\n" + context.GetObjectGraphString());
            }

            return ReflectionUtil.CreateGenericList(context.MemberType, new object[] {});
        }

        public List<Type> ResolveTypeAll(InjectContext context)
        {
            if (_providers.ContainsKey(context.BindingId))
            {
                return _providers[context.BindingId].Select(x => x.GetInstanceType()).Where(x => x != null).ToList();
            }

            return new List<Type> {};
        }

        public void Install(IEnumerable<IInstaller> installers)
        {
            foreach (var installer in installers)
            {
                Assert.IsNotNull(installer, "Tried to install a null installer");

                if (installer.IsEnabled)
                {
                    Install(installer);
                }
            }
        }

        public void Install(IInstaller installer)
        {
            Assert.That(installer.IsEnabled);

            this.Inject(installer);
            InstallInstallerInternal(installer);
        }

        public void Install<T>(params object[] extraArgs)
            where T : IInstaller
        {
            Install(typeof(T), extraArgs);
        }

        public void Install(Type installerType, params object[] extraArgs)
        {
            Assert.That(installerType.DerivesFrom<IInstaller>());

#if !ZEN_NOT_UNITY3D
            if (installerType.DerivesFrom<MonoInstaller>())
            {
                var installer = InstantiatePrefabResourceForComponent<MonoInstaller>("Installers/" + installerType.Name(), extraArgs);

                try
                {
                    InstallInstallerInternal(installer);
                }
                finally
                {
                    // When running, it is nice to keep the installer around so that you can change the settings live
                    // But when validating at edit time, we don't want to add the new game object
                    if (!Application.isPlaying)
                    {
                        GameObject.DestroyImmediate(installer.gameObject);
                    }
                }
            }
            else
#endif
            {
                var installer = (IInstaller)this.Instantiate(installerType, extraArgs);
                InstallInstallerInternal(installer);
            }
        }

        public bool HasInstalled<T>()
            where T : IInstaller
        {
            return HasInstalled(typeof(T));
        }

        public bool HasInstalled(Type installerType)
        {
            return _installedInstallers.Where(x => x == installerType).Any();
        }

        void InstallInstallerInternal(IInstaller installer)
        {
            var installerType = installer.GetType();

            Log.Debug("Installing installer '{0}'", installerType);

            Assert.That(!_installsInProgress.Contains(installerType),
                "Potential infinite loop detected while installing '{0}'", installerType.Name());

            Assert.That(!_installedInstallers.Contains(installerType),
                "Tried installing installer '{0}' twice", installerType.Name());

            _installedInstallers.Add(installerType);
            _installsInProgress.Push(installerType);

            try
            {
                installer.InstallBindings();
            }
            catch (Exception e)
            {
                // This context information is really helpful when bind commands fail
                throw new Exception(
                    "Error occurred while running installer '{0}'".Fmt(installer.GetType().Name()), e);
            }
            finally
            {
                Assert.That(_installsInProgress.Peek().Equals(installerType));
                _installsInProgress.Pop();
            }
        }

        // Try looking up a single provider for a given context
        // Note that this method should not throw zenject exceptions
        internal ProviderLookupResult TryGetUniqueProvider(
            InjectContext context, out ProviderBase provider)
        {
            // Note that different types can map to the same provider (eg. a base type to a concrete class and a concrete class to itself)
            var providers = GetProviderMatchesInternal(context).ToList();

            if (providers.IsEmpty())
            {
                provider = null;
                return ProviderLookupResult.None;
            }

            if (providers.Count > 1)
            {
                // If we find multiple providers and we are looking for just one, then
                // try to intelligently choose one from the list before giving up

                // First try picking the most 'local' dependencies
                // This will bias towards bindings for the lower level specific containers rather than the global high level container
                // This will, for example, allow you to just ask for a DiContainer dependency without needing to specify [InjectLocal]
                // (otherwise it would always match for a list of DiContainer's for all parent containers)
                var sortedProviders = providers.Select(x => new { Pair = x, Distance = GetContainerHeirarchyDistance(x.Container) }).OrderBy(x => x.Distance).ToList();

                sortedProviders.RemoveAll(x => x.Distance != sortedProviders[0].Distance);

                if (sortedProviders.Count == 1)
                {
                    // We have one match that is the closest
                    provider = sortedProviders[0].Pair.Provider;
                }
                else
                {
                    // Try choosing the one with a condition before giving up and throwing an exception
                    // This is nice because it allows us to bind a default and then override with conditions
                    provider = sortedProviders.Select(x => x.Pair.Provider).Where(x => x.Condition != null).OnlyOrDefault();

                    if (provider == null)
                    {
                        return ProviderLookupResult.Multiple;
                    }
                }
            }
            else
            {
                provider = providers.Single().Provider;
            }

            Assert.IsNotNull(provider);
            return ProviderLookupResult.Success;
        }

        // Return single instance of requested type or assert
        public object Resolve(InjectContext context)
        {
            ProviderBase provider;

            var result = TryGetUniqueProvider(context, out provider);

            if (result == ProviderLookupResult.Multiple)
            {
                throw new ZenjectResolveException(
                    "Found multiple matches when only one was expected for type '{0}'{1}. \nObject graph:\n {2}"
                    .Fmt(
                        context.MemberType.Name(),
                        (context.ObjectType == null ? "" : " while building object with type '{0}'".Fmt(context.ObjectType.Name())),
                        context.GetObjectGraphString()));
            }

            if (result == ProviderLookupResult.None)
            {
                // If it's a generic list then try matching multiple instances to its generic type
                if (ReflectionUtil.IsGenericList(context.MemberType))
                {
                    var subType = context.MemberType.GetGenericArguments().Single();
                    var subContext = context.ChangeMemberType(subType);

                    return ResolveAll(subContext);
                }

                if (context.Optional)
                {
                    return context.FallBackValue;
                }

                throw new ZenjectResolveException(
                    "Unable to resolve type '{0}'{1}. \nObject graph:\n{2}"
                    .Fmt(
                        context.MemberType.Name() + (context.Identifier == null ? "" : " with ID '" + context.Identifier.ToString() + "'"),
                        (context.ObjectType == null ? "" : " while building object with type '{0}'".Fmt(context.ObjectType.Name())),
                        context.GetObjectGraphString()));
            }

            Assert.That(result == ProviderLookupResult.Success);
            Assert.IsNotNull(provider);

            return SafeGetInstance(provider, context);
        }

        object SafeGetInstance(ProviderBase provider, InjectContext context)
        {
            if (ChecksForCircularDependencies)
            {
                var lookupId = new LookupId(provider, context.BindingId);

                // Allow one before giving up so that you can do circular dependencies via postinject or fields
                if (_resolvesInProgress.Where(x => x.Equals(lookupId)).Count() > 1)
                {
                    throw new ZenjectResolveException(
                        "Circular dependency detected! \nObject graph:\n {0}".Fmt(context.GetObjectGraphString()));
                }

                _resolvesInProgress.Push(lookupId);
                try
                {
                    return provider.GetInstance(context);
                }
                finally
                {
                    Assert.That(_resolvesInProgress.Peek().Equals(lookupId));
                    _resolvesInProgress.Pop();
                }
            }
            else
            {
                return provider.GetInstance(context);
            }
        }

        int GetContainerHeirarchyDistance(DiContainer container)
        {
            return GetContainerHeirarchyDistance(container, 0);
        }

        int GetContainerHeirarchyDistance(DiContainer container, int depth)
        {
            if (container == this)
            {
                return depth;
            }

            Assert.IsNotNull(_parentContainer);
            return _parentContainer.GetContainerHeirarchyDistance(container, depth + 1);
        }

        public IEnumerable<Type> GetDependencyContracts(Type contract)
        {
            foreach (var injectMember in TypeAnalyzer.GetInfo(contract).AllInjectables)
            {
                yield return injectMember.MemberType;
            }
        }

        // Same as Instantiate except you can pass in null value
        // however the type for each parameter needs to be explicitly provided in this case
        public object InstantiateExplicit(
            Type concreteType, List<TypeValuePair> extraArgMap, InjectContext currentContext, string concreteIdentifier, bool autoInject)
        {
#if PROFILING_ENABLED
            using (ProfileBlock.Start("Zenject.Instantiate({0})", concreteType))
#endif
            {
                return InstantiateInternal(concreteType, extraArgMap, currentContext, concreteIdentifier, autoInject);
            }
        }

        object InstantiateInternal(
            Type concreteType, IEnumerable<TypeValuePair> extraArgs, InjectContext currentContext, string concreteIdentifier, bool autoInject)
        {
#if !ZEN_NOT_UNITY3D
            Assert.That(!concreteType.DerivesFrom<UnityEngine.Component>(),
                "Error occurred while instantiating object of type '{0}'. Instantiator should not be used to create new mono behaviours.  Must use InstantiatePrefabForComponent, InstantiatePrefab, InstantiateComponentOnNewGameObject, InstantiateGameObject, or InstantiateComponent.  You may also want to use GameObjectFactory class or plain old GameObject.Instantiate.", concreteType.Name());
#endif

            var typeInfo = TypeAnalyzer.GetInfo(concreteType);

            if (typeInfo.InjectConstructor == null)
            {
                throw new ZenjectResolveException(
                    "More than one (or zero) constructors found for type '{0}' when creating dependencies.  Use one [Inject] attribute to specify which to use.".Fmt(concreteType));
            }

            // Make a copy since we remove from it below
            var extraArgList = extraArgs.ToList();
            var paramValues = new List<object>();

            foreach (var injectInfo in typeInfo.ConstructorInjectables)
            {
                object value;

                if (!InstantiateUtil.PopValueWithType(extraArgList, injectInfo.MemberType, out value))
                {
                    value = Resolve(injectInfo.CreateInjectContext(this, currentContext, null, concreteIdentifier));
                }

                paramValues.Add(value);
            }

            object newObj;

            //Log.Debug("Zenject: Instantiating type '{0}'", concreteType.Name());
            try
            {
#if PROFILING_ENABLED
                using (ProfileBlock.Start("{0}.{0}()", concreteType))
#endif
                {
                    newObj = typeInfo.InjectConstructor.Invoke(paramValues.ToArray());
                }
            }
            catch (Exception e)
            {
                throw new ZenjectResolveException(
                    "Error occurred while instantiating object with type '{0}'".Fmt(concreteType.Name()), e);
            }

            if (autoInject)
            {
                InjectExplicit(newObj, extraArgList, true, typeInfo, currentContext, concreteIdentifier);
            }
            else
            {
                if (!extraArgList.IsEmpty())
                {
                    throw new ZenjectResolveException(
                        "Passed unnecessary parameters when injecting into type '{0}'. \nExtra Parameters: {1}\nObject graph:\n{2}"
                        .Fmt(newObj.GetType().Name(), String.Join(",", extraArgList.Select(x => x.Type.Name()).ToArray()), currentContext.GetObjectGraphString()));
                }
            }

            return newObj;
        }

        // Iterate over fields/properties on the given object and inject any with the [Inject] attribute
        public void InjectExplicit(
            object injectable, IEnumerable<TypeValuePair> extraArgs,
            bool shouldUseAll, ZenjectTypeInfo typeInfo, InjectContext context, string concreteIdentifier)
        {
            Assert.IsEqual(typeInfo.TypeAnalyzed, injectable.GetType());
            Assert.That(injectable != null);

#if !ZEN_NOT_UNITY3D
            Assert.That(injectable.GetType() != typeof(GameObject),
                "Use InjectGameObject to Inject game objects instead of Inject method");
#endif

            // Make a copy since we remove from it below
            var extraArgsList = extraArgs.ToList();

            foreach (var injectInfo in typeInfo.FieldInjectables.Concat(typeInfo.PropertyInjectables))
            {
                object value;

                if (InstantiateUtil.PopValueWithType(extraArgsList, injectInfo.MemberType, out value))
                {
                    injectInfo.Setter(injectable, value);
                }
                else
                {
                    value = Resolve(
                        injectInfo.CreateInjectContext(this, context, injectable, concreteIdentifier));

                    if (injectInfo.Optional && value == null)
                    {
                        // Do not override in this case so it retains the hard-coded value
                    }
                    else
                    {
                        injectInfo.Setter(injectable, value);
                    }
                }
            }

            foreach (var method in typeInfo.PostInjectMethods)
            {
#if PROFILING_ENABLED
                using (ProfileBlock.Start("{0}.{1}()", injectable.GetType(), method.MethodInfo.Name))
#endif
                {
                    var paramValues = new List<object>();

                    foreach (var injectInfo in method.InjectableInfo)
                    {
                        object value;

                        if (!InstantiateUtil.PopValueWithType(extraArgsList, injectInfo.MemberType, out value))
                        {
                            value = Resolve(
                                injectInfo.CreateInjectContext(this, context, injectable, concreteIdentifier));
                        }

                        paramValues.Add(value);
                    }

                    method.MethodInfo.Invoke(injectable, paramValues.ToArray());
                }
            }

            if (shouldUseAll && !extraArgsList.IsEmpty())
            {
                throw new ZenjectResolveException(
                    "Passed unnecessary parameters when injecting into type '{0}'. \nExtra Parameters: {1}\nObject graph:\n{2}"
                    .Fmt(injectable.GetType().Name(), String.Join(",", extraArgsList.Select(x => x.Type.Name()).ToArray()), context.GetObjectGraphString()));
            }
        }

#if !ZEN_NOT_UNITY3D

        // NOTE: gameobject here is not a prefab prototype, it is an instance
        public Component InstantiateComponent(
            Type componentType, GameObject gameObject, params object[] extraArgMap)
        {
            Assert.That(componentType.DerivesFrom<Component>());

            var monoBehaviour = (Component)gameObject.AddComponent(componentType);
            this.Inject(monoBehaviour, extraArgMap);
            return monoBehaviour;
        }

        public GameObject InstantiatePrefabResourceExplicit(
            string resourcePath, IEnumerable<object> extraArgMap, InjectContext context)
        {
            return InstantiatePrefabResourceExplicit(resourcePath, extraArgMap, context, false);
        }

        public GameObject InstantiatePrefabResourceExplicit(
            string resourcePath, IEnumerable<object> extraArgMap, InjectContext context, bool includeInactive)
        {
            var prefab = (GameObject)Resources.Load(resourcePath);
            Assert.IsNotNull(prefab, "Could not find prefab at resource location '{0}'".Fmt(resourcePath));
            return InstantiatePrefabExplicit(prefab, extraArgMap, context, includeInactive);
        }

        public GameObject InstantiatePrefabExplicit(
            GameObject prefab, IEnumerable<object> extraArgMap, InjectContext context)
        {
            return InstantiatePrefabExplicit(prefab, extraArgMap, context, false);
        }

        public GameObject InstantiatePrefabExplicit(
            GameObject prefab, IEnumerable<object> extraArgMap, InjectContext context, bool includeInactive)
        {
            var gameObj = (GameObject)GameObject.Instantiate(prefab);

            if (_rootTransform != null)
            {
                // By default parent to comp root
                // This is good so that the entire object graph is
                // contained underneath it, which is useful for cases
                // where you need to delete the entire object graph
                gameObj.transform.SetParent(_rootTransform, false);
            }

            gameObj.SetActive(true);

            this.InjectGameObject(gameObj, true, includeInactive, extraArgMap, context);

            return gameObj;
        }

        // Create a new empty game object under the composition root
        public GameObject InstantiateGameObject(string name)
        {
            var gameObj = new GameObject(name);

            if (_rootTransform != null)
            {
                gameObj.transform.SetParent(_rootTransform, false);
            }

            return gameObj;
        }

        public object InstantiateComponentOnNewGameObjectExplicit(
            Type componentType, string name, List<TypeValuePair> extraArgMap, InjectContext currentContext)
        {
            Assert.That(componentType.DerivesFrom<Component>(), "Expected type '{0}' to derive from UnityEngine.Component", componentType.Name());

            var gameObj = new GameObject(name);

            if (_rootTransform != null)
            {
                gameObj.transform.SetParent(_rootTransform, false);
            }

            if (componentType == typeof(Transform))
            {
                Assert.That(extraArgMap.IsEmpty());
                return gameObj.transform;
            }

            var component = (Component)gameObj.AddComponent(componentType);

            this.InjectExplicit(component, extraArgMap, currentContext);

            return component;
        }

        public object InstantiatePrefabResourceForComponentExplicit(
            Type componentType, string resourcePath, List<TypeValuePair> extraArgs, InjectContext currentContext)
        {
            var prefab = (GameObject)Resources.Load(resourcePath);
            Assert.IsNotNull(prefab, "Could not find prefab at resource location '{0}'".Fmt(resourcePath));
            return InstantiatePrefabForComponentExplicit(
                componentType, prefab, extraArgs, currentContext);
        }

        public object InstantiatePrefabForComponentExplicit(
            Type componentType, GameObject prefab, List<TypeValuePair> extraArgs, InjectContext currentContext)
        {
            return InstantiatePrefabForComponentExplicit(componentType, prefab, extraArgs, currentContext, false);
        }

        public object InstantiatePrefabForComponentExplicit(
            Type componentType, GameObject prefab, List<TypeValuePair> extraArgs, InjectContext currentContext, bool includeInactive)
        {
            Assert.That(prefab != null, "Null prefab found when instantiating game object");

            // It could be an interface so this may fail in valid cases so you may want to comment out
            // Leaving it in for now to catch the more likely scenario of it being a mistake
            Assert.That(componentType.DerivesFrom<Component>(), "Expected type '{0}' to derive from UnityEngine.Component", componentType.Name());

            var gameObj = (GameObject)GameObject.Instantiate(prefab);

            if (_rootTransform != null)
            {
                // By default parent to comp root
                // This is good so that the entire object graph is
                // contained underneath it, which is useful for cases
                // where you need to delete the entire object graph
                gameObj.transform.SetParent(_rootTransform, false);
            }

            gameObj.SetActive(true);

            Component requestedScript = null;

            // Inject on the children first since the parent objects are more likely to use them in their post inject methods
            foreach (var component in UnityUtil.GetComponentsInChildrenBottomUp(gameObj, includeInactive))
            {
                if (component != null)
                {
                    if (component.GetType().DerivesFromOrEqual(componentType))
                    {
                        Assert.IsNull(requestedScript,
                            "Found multiple matches with type '{0}' when instantiating new game object from prefab '{1}'", componentType, prefab.name);
                        requestedScript = component;

                        this.InjectExplicit(component, extraArgs);
                    }
                    else
                    {
                        this.Inject(component);
                    }
                }
                else
                {
                    Log.Warn("Found null component while instantiating prefab '{0}'.  Possible missing script.", prefab.name);
                }
            }

            if (requestedScript == null)
            {
                throw new ZenjectResolveException(
                    "Could not find component with type '{0}' when instantiating new game object".Fmt(componentType));
            }

            return requestedScript;
        }
#endif

        ////////////// Convenience methods for IInstantiator ////////////////

        public T Instantiate<T>(
            params object[] extraArgs)
        {
            return (T)Instantiate(typeof(T), extraArgs);
        }

        public object Instantiate(
            Type concreteType, params object[] extraArgs)
        {
            Assert.That(!extraArgs.ContainsItem(null),
                "Null value given to factory constructor arguments when instantiating object with type '{0}'. In order to use null use InstantiateExplicit", concreteType);

            return InstantiateExplicit(
                concreteType, InstantiateUtil.CreateTypeValueList(extraArgs));
        }

        // This is used instead of Instantiate to support specifying null values
        public T InstantiateExplicit<T>(
            List<TypeValuePair> extraArgMap)
        {
            return (T)InstantiateExplicit(typeof(T), extraArgMap);
        }

        public T InstantiateExplicit<T>(
            List<TypeValuePair> extraArgMap, InjectContext context)
        {
            return (T)InstantiateExplicit(
                typeof(T), extraArgMap, context);
        }

        public object InstantiateExplicit(
            Type concreteType, List<TypeValuePair> extraArgMap)
        {
            return InstantiateExplicit(
                concreteType, extraArgMap, new InjectContext(this, concreteType, null));
        }

        public object InstantiateExplicit(
            Type concreteType, List<TypeValuePair> extraArgMap, InjectContext context)
        {
            return InstantiateExplicit(
                concreteType, extraArgMap, context, null, true);
        }

#if !ZEN_NOT_UNITY3D
        public TContract InstantiateComponent<TContract>(
            GameObject gameObject, params object[] args)
            where TContract : Component
        {
            return (TContract)InstantiateComponent(typeof(TContract), gameObject, args);
        }

        public GameObject InstantiatePrefab(
            GameObject prefab, params object[] args)
        {
            return InstantiatePrefabExplicit(prefab, args, null);
        }

        public GameObject InstantiatePrefab(
            bool includeInactive, GameObject prefab, params object[] args)
        {
            return InstantiatePrefabExplicit(prefab, args, null, includeInactive);
        }

        public GameObject InstantiatePrefabResource(
            string resourcePath, params object[] args)
        {
            return InstantiatePrefabResourceExplicit(resourcePath, args, null, false);
        }

        public GameObject InstantiatePrefabResource(
            bool includeInactive, string resourcePath, params object[] args)
        {
            return InstantiatePrefabResourceExplicit(resourcePath, args, null, includeInactive);
        }

        /////////////// InstantiatePrefabForComponent

        public T InstantiatePrefabForComponent<T>(
            GameObject prefab, params object[] extraArgs)
        {
            return (T)InstantiatePrefabForComponent(typeof(T), prefab, extraArgs);
        }

        public object InstantiatePrefabForComponent(
            Type concreteType, GameObject prefab, params object[] extraArgs)
        {
            Assert.That(!extraArgs.ContainsItem(null),
                "Null value given to factory constructor arguments when instantiating object with type '{0}'. In order to use null use InstantiatePrefabForComponentExplicit", concreteType);

            return InstantiatePrefabForComponentExplicit(
                concreteType, prefab, InstantiateUtil.CreateTypeValueList(extraArgs));
        }

        public T InstantiatePrefabForComponent<T>(
            bool includeInactive, GameObject prefab, params object[] extraArgs)
        {
            return (T)InstantiatePrefabForComponent(includeInactive, typeof(T), prefab, extraArgs);
        }

        public object InstantiatePrefabForComponent(
            bool includeInactive, Type concreteType, GameObject prefab, params object[] extraArgs)
        {
            Assert.That(!extraArgs.Contains(null),
            "Null value given to factory constructor arguments when instantiating object with type '{0}'. In order to use null use InstantiatePrefabForComponentExplicit", concreteType);

            return InstantiatePrefabForComponentExplicit(
                concreteType, prefab,
                InstantiateUtil.CreateTypeValueList(extraArgs),
                new InjectContext(this, concreteType, null), includeInactive);
        }

        // This is used instead of Instantiate to support specifying null values
        public T InstantiatePrefabForComponentExplicit<T>(
            GameObject prefab, List<TypeValuePair> extraArgMap)
        {
            return (T)InstantiatePrefabForComponentExplicit(typeof(T), prefab, extraArgMap);
        }

        public object InstantiatePrefabForComponentExplicit(
            Type concreteType, GameObject prefab, List<TypeValuePair> extraArgMap)
        {
            return InstantiatePrefabForComponentExplicit(
                concreteType, prefab, extraArgMap, new InjectContext(this, concreteType, null));
        }


        /////////////// InstantiatePrefabForComponent

        public T InstantiatePrefabResourceForComponent<T>(
            string resourcePath, params object[] extraArgs)
        {
            return (T)InstantiatePrefabResourceForComponent(typeof(T), resourcePath, extraArgs);
        }

        public object InstantiatePrefabResourceForComponent(
            Type concreteType, string resourcePath, params object[] extraArgs)
        {
            Assert.That(!extraArgs.ContainsItem(null),
            "Null value given to factory constructor arguments when instantiating object with type '{0}'. In order to use null use InstantiatePrefabForComponentExplicit", concreteType);

            return InstantiatePrefabResourceForComponentExplicit(
                concreteType, resourcePath, InstantiateUtil.CreateTypeValueList(extraArgs));
        }

        // This is used instead of Instantiate to support specifying null values
        public T InstantiatePrefabResourceForComponentExplicit<T>(
            string resourcePath, List<TypeValuePair> extraArgMap)
        {
            return (T)InstantiatePrefabResourceForComponentExplicit(typeof(T), resourcePath, extraArgMap);
        }

        public object InstantiatePrefabResourceForComponentExplicit(
            Type concreteType, string resourcePath, List<TypeValuePair> extraArgMap)
        {
            return InstantiatePrefabResourceForComponentExplicit(
                concreteType, resourcePath, extraArgMap, new InjectContext(this, concreteType, null));
        }

        /////////////// InstantiateComponentOnNewGameObject

        public T InstantiateComponentOnNewGameObject<T>(
            string name, params object[] extraArgs)
        {
            return (T)InstantiateComponentOnNewGameObject(typeof(T), name, extraArgs);
        }

        public object InstantiateComponentOnNewGameObject(
            Type concreteType, string name, params object[] extraArgs)
        {
            Assert.That(!extraArgs.ContainsItem(null),
                "Null value given to factory constructor arguments when instantiating object with type '{0}'. In order to use null use InstantiateComponentOnNewGameObjectExplicit", concreteType);

            return InstantiateComponentOnNewGameObjectExplicit(
                concreteType, name, InstantiateUtil.CreateTypeValueList(extraArgs));
        }

        // This is used instead of Instantiate to support specifying null values
        public T InstantiateComponentOnNewGameObjectExplicit<T>(
            string name, List<TypeValuePair> extraArgMap)
        {
            return (T)InstantiateComponentOnNewGameObjectExplicit(typeof(T), name, extraArgMap);
        }

        public object InstantiateComponentOnNewGameObjectExplicit(
            Type concreteType, string name, List<TypeValuePair> extraArgMap)
        {
            return InstantiateComponentOnNewGameObjectExplicit(
                concreteType, name, extraArgMap, new InjectContext(this, concreteType, null));
        }
#endif

        ////////////// Convenience methods for IResolver ////////////////

#if !ZEN_NOT_UNITY3D
        // Inject dependencies into child game objects
        public void InjectGameObject(
            GameObject gameObject, bool recursive, bool includeInactive)
        {
            InjectGameObject(gameObject, recursive, includeInactive, Enumerable.Empty<object>());
        }

        public void InjectGameObject(
            GameObject gameObject, bool recursive)
        {
            InjectGameObject(gameObject, recursive, false);
        }

        public void InjectGameObject(
            GameObject gameObject)
        {
            InjectGameObject(gameObject, true, false);
        }

        public void InjectGameObject(
            GameObject gameObject,
            bool recursive, bool includeInactive, IEnumerable<object> extraArgs)
        {
            InjectGameObject(
                gameObject, recursive, includeInactive, extraArgs, null);
        }

        public void InjectGameObject(
            GameObject gameObject,
            bool recursive, bool includeInactive, IEnumerable<object> extraArgs, InjectContext context)
        {
            IEnumerable<Component> components;

            if (recursive)
            {
                components = UnityUtil.GetComponentsInChildrenBottomUp(gameObject, includeInactive);
            }
            else
            {
                if (!includeInactive && !gameObject.activeSelf)
                {
                    return;
                }

                components = gameObject.GetComponents<Component>();
            }

            foreach (var component in components)
            {
                // null if monobehaviour link is broken
                // Do not inject on installers since these are always injected before they are installed
                if (component != null && !component.GetType().DerivesFrom<MonoInstaller>())
                {
                    Inject(component, extraArgs, false, context);
                }
            }
        }
#endif

        public void Inject(object injectable)
        {
            Inject(injectable, Enumerable.Empty<object>());
        }

        public void Inject(object injectable, IEnumerable<object> additional)
        {
            Inject(injectable, additional, true);
        }

        public void Inject(object injectable, IEnumerable<object> additional, bool shouldUseAll)
        {
            Inject(
                injectable, additional, shouldUseAll, new InjectContext(this, injectable.GetType(), null));
        }

        public void Inject(
            object injectable, IEnumerable<object> additional, bool shouldUseAll, InjectContext context)
        {
            Inject(
                injectable, additional, shouldUseAll, context, TypeAnalyzer.GetInfo(injectable.GetType()));
        }

        public void Inject(
            object injectable,
            IEnumerable<object> additional, bool shouldUseAll, InjectContext context, ZenjectTypeInfo typeInfo)
        {
            Assert.That(!additional.ContainsItem(null),
                "Null value given to injection argument list. In order to use null you must provide a List<TypeValuePair> and not just a list of objects");

            InjectExplicit(
                injectable, InstantiateUtil.CreateTypeValueList(additional), shouldUseAll, typeInfo, context, null);
        }

        public void InjectExplicit(object injectable, List<TypeValuePair> additional)
        {
            InjectExplicit(
                injectable, additional, new InjectContext(this, injectable.GetType(), null));
        }

        public void InjectExplicit(object injectable, List<TypeValuePair> additional, InjectContext context)
        {
            InjectExplicit(
                injectable, additional, true,
                TypeAnalyzer.GetInfo(injectable.GetType()), context, null);
        }

        public List<Type> ResolveTypeAll(Type type)
        {
            return ResolveTypeAll(new InjectContext(this, type, null));
        }

        public TContract Resolve<TContract>()
        {
            return Resolve<TContract>((string)null);
        }

        public TContract Resolve<TContract>(string identifier)
        {
            return Resolve<TContract>(new InjectContext(this, typeof(TContract), identifier));
        }

        public TContract TryResolve<TContract>()
            where TContract : class
        {
            return TryResolve<TContract>((string)null);
        }

        public TContract TryResolve<TContract>(string identifier)
            where TContract : class
        {
            return (TContract)TryResolve(typeof(TContract), identifier);
        }

        public object TryResolve(Type contractType)
        {
            return TryResolve(contractType, null);
        }

        public object TryResolve(Type contractType, string identifier)
        {
            return Resolve(new InjectContext(this, contractType, identifier, true));
        }

        public object Resolve(Type contractType)
        {
            return Resolve(new InjectContext(this, contractType, null));
        }

        public object Resolve(Type contractType, string identifier)
        {
            return Resolve(new InjectContext(this, contractType, identifier));
        }

        public TContract Resolve<TContract>(InjectContext context)
        {
            Assert.IsEqual(context.MemberType, typeof(TContract));
            return (TContract) Resolve(context);
        }

        public List<TContract> ResolveAll<TContract>()
        {
            return ResolveAll<TContract>((string)null);
        }

        public List<TContract> ResolveAll<TContract>(bool optional)
        {
            return ResolveAll<TContract>(null, optional);
        }

        public List<TContract> ResolveAll<TContract>(string identifier)
        {
            return ResolveAll<TContract>(identifier, false);
        }

        public List<TContract> ResolveAll<TContract>(string identifier, bool optional)
        {
            var context = new InjectContext(this, typeof(TContract), identifier, optional);
            return ResolveAll<TContract>(context);
        }

        public List<TContract> ResolveAll<TContract>(InjectContext context)
        {
            Assert.IsEqual(context.MemberType, typeof(TContract));
            return (List<TContract>) ResolveAll(context);
        }

        public IList ResolveAll(Type contractType)
        {
            return ResolveAll(contractType, null);
        }

        public IList ResolveAll(Type contractType, string identifier)
        {
            return ResolveAll(contractType, identifier, false);
        }

        public IList ResolveAll(Type contractType, bool optional)
        {
            return ResolveAll(contractType, null, optional);
        }

        public IList ResolveAll(Type contractType, string identifier, bool optional)
        {
            var context = new InjectContext(this, contractType, identifier, optional);
            return ResolveAll(context);
        }

        ////////////// IBinder ////////////////

        public bool Unbind<TContract>(string identifier)
        {
            List<ProviderBase> providersToRemove;
            var bindingId = new BindingId(typeof(TContract), identifier);

            if (_providers.TryGetValue(bindingId, out providersToRemove))
            {
                _providers.Remove(bindingId);

                // Only dispose if the provider is not bound to another type
                foreach (var provider in providersToRemove)
                {
                    if (_providers.Where(x => x.Value.ContainsItem(provider)).IsEmpty())
                    {
                        provider.Dispose();
                    }
                }

                return true;
            }

            return false;
        }

        public BindingConditionSetter BindInstance<TContract>(string identifier, TContract obj)
        {
            return Bind<TContract>(identifier).ToInstance(obj);
        }

        public BindingConditionSetter BindInstance<TContract>(TContract obj)
        {
            return Bind<TContract>().ToInstance(obj);
        }

        public GenericBinder<TContract> Bind<TContract>()
        {
            return Bind<TContract>(null);
        }

        public UntypedBinder Bind(Type contractType)
        {
            return Bind(contractType, null);
        }

        public bool Unbind<TContract>()
        {
            return Unbind<TContract>(null);
        }

        public bool HasBinding(InjectContext context)
        {
            List<ProviderBase> providers;

            if (!_providers.TryGetValue(context.BindingId, out providers))
            {
                return false;
            }

            return providers.Where(x => x.Matches(context)).HasAtLeast(1);
        }

        public bool HasBinding<TContract>()
        {
            return HasBinding<TContract>(null);
        }

        public bool HasBinding<TContract>(string identifier)
        {
            return HasBinding(
                new InjectContext(this, typeof(TContract), identifier));
        }

        public void BindAllInterfacesToSingle<TConcrete>()
        {
            BindAllInterfacesToSingle(typeof(TConcrete));
        }

        public void BindAllInterfacesToSingle(Type concreteType)
        {
            foreach (var interfaceType in concreteType.GetInterfaces())
            {
                Assert.That(concreteType.DerivesFrom(interfaceType));
                Bind(interfaceType).ToSingle(concreteType);
            }
        }

        public void BindAllInterfacesToInstance(object value)
        {
            BindAllInterfacesToInstance(value.GetType(), value);
        }

        public void BindAllInterfacesToInstance(Type concreteType, object value)
        {
            Assert.That((value == null && IsValidating) || value.GetType().DerivesFromOrEqual(concreteType));

            foreach (var interfaceType in concreteType.GetInterfaces())
            {
                Assert.That(concreteType.DerivesFrom(interfaceType));
                Bind(interfaceType).ToInstance(concreteType, value);
            }
        }

        public IFactoryUntypedBinder<TContract> BindIFactoryUntyped<TContract>(string identifier)
        {
            return new IFactoryUntypedBinder<TContract>(this, identifier);
        }

        public IFactoryUntypedBinder<TContract> BindIFactoryUntyped<TContract>()
        {
            return BindIFactoryUntyped<TContract>(null);
        }

        public IFactoryBinder<TContract> BindIFactory<TContract>(string identifier)
        {
            return new IFactoryBinder<TContract>(this, identifier);
        }

        public IFactoryBinder<TContract> BindIFactory<TContract>()
        {
            return BindIFactory<TContract>(null);
        }

        public IFactoryBinder<TParam1, TContract> BindIFactory<TParam1, TContract>(string identifier)
        {
            return new IFactoryBinder<TParam1, TContract>(this, identifier);
        }

        public IFactoryBinder<TParam1, TContract> BindIFactory<TParam1, TContract>()
        {
            return BindIFactory<TParam1, TContract>(null);
        }

        public IFactoryBinder<TParam1, TParam2, TContract> BindIFactory<TParam1, TParam2, TContract>(string identifier)
        {
            return new IFactoryBinder<TParam1, TParam2, TContract>(this, identifier);
        }

        public IFactoryBinder<TParam1, TParam2, TContract> BindIFactory<TParam1, TParam2, TContract>()
        {
            return BindIFactory<TParam1, TParam2, TContract>(null);
        }

        public IFactoryBinder<TParam1, TParam2, TParam3, TContract> BindIFactory<TParam1, TParam2, TParam3, TContract>(string identifier)
        {
            return new IFactoryBinder<TParam1, TParam2, TParam3, TContract>(this, identifier);
        }

        public IFactoryBinder<TParam1, TParam2, TParam3, TContract> BindIFactory<TParam1, TParam2, TParam3, TContract>()
        {
            return BindIFactory<TParam1, TParam2, TParam3, TContract>(null);
        }

        public IFactoryBinder<TParam1, TParam2, TParam3, TParam4, TContract> BindIFactory<TParam1, TParam2, TParam3, TParam4, TContract>(string identifier)
        {
            return new IFactoryBinder<TParam1, TParam2, TParam3, TParam4, TContract>(this, identifier);
        }

        public IFactoryBinder<TParam1, TParam2, TParam3, TParam4, TContract> BindIFactory<TParam1, TParam2, TParam3, TParam4, TContract>()
        {
            return BindIFactory<TParam1, TParam2, TParam3, TParam4, TContract>(null);
        }

        public GenericBinder<TContract> Rebind<TContract>()
        {
            this.Unbind<TContract>();
            return this.Bind<TContract>();
        }

        public GenericBinder<TContract> Bind<TContract>(string identifier)
        {
            Assert.That(!typeof(TContract).DerivesFromOrEqual<IInstaller>(),
                "Deprecated usage of Bind<IInstaller>, use Install<IInstaller> instead");
            return new GenericBinder<TContract>(this, identifier, _singletonMap);
        }

        public FacadeBinder<TFacade> BindFacade<TFacade>(Action<DiContainer> installerFunc)
            where TFacade : IFacade
        {
            return BindFacade<TFacade>(installerFunc, null);
        }

        public FacadeBinder<TFacade> BindFacade<TFacade>(
            Action<DiContainer> installerFunc, string identifier)
            where TFacade : IFacade
        {
            return new FacadeBinder<TFacade>(this, identifier, installerFunc);
        }

        // Note that this can include open generic types as well such as List<>
        public UntypedBinder Bind(Type contractType, string identifier)
        {
            Assert.That(!contractType.DerivesFromOrEqual<IInstaller>(),
                "Deprecated usage of Bind<IInstaller>, use Install<IInstaller> instead");
            return new UntypedBinder(this, contractType, identifier, _singletonMap);
        }

#if !ZEN_NOT_UNITY3D
        public BindingConditionSetter BindGameObjectFactory<T>(
            GameObject prefab)
            // This would be useful but fails with VerificationException's in webplayer builds for some reason
            //where T : GameObjectFactory
            where T : class
        {
            if (prefab == null)
            {
                throw new ZenjectBindException(
                    "Null prefab provided to BindGameObjectFactory for type '{0}'".Fmt(typeof(T).Name()));
            }

            // We could bind the factory ToSingle but doing it this way is better
            // since it allows us to have multiple game object factories that
            // use different prefabs and have them injected into different places
            return Bind<T>().ToMethod((ctx) => ctx.Container.Instantiate<T>(prefab));
        }
#endif

        public BindingConditionSetter BindFacadeFactory<TFacade, TFacadeFactory>(
            Action<DiContainer> facadeInstaller)
            where TFacade : IFacade
            where TFacadeFactory : FacadeFactory<TFacade>
        {
            return this.Bind<TFacadeFactory>().ToMethod(
                x => x.Container.Instantiate<TFacadeFactory>(facadeInstaller));
        }

        public BindingConditionSetter BindFacadeFactory<TParam1, TFacade, TFacadeFactory>(
            Action<DiContainer, TParam1> facadeInstaller)
            where TFacade : IFacade
            where TFacadeFactory : FacadeFactory<TParam1, TFacade>
        {
            return this.Bind<TFacadeFactory>().ToMethod(
                x => x.Container.Instantiate<TFacadeFactory>(facadeInstaller));
        }

        public BindingConditionSetter BindFacadeFactory<TParam1, TParam2, TFacade, TFacadeFactory>(
            Action<DiContainer, TParam1, TParam2> facadeInstaller)
            where TFacade : IFacade
            where TFacadeFactory : FacadeFactory<TParam1, TParam2, TFacade>
        {
            return this.Bind<TFacadeFactory>().ToMethod(
                x => x.Container.Instantiate<TFacadeFactory>(facadeInstaller));
        }

        public BindingConditionSetter BindFacadeFactory<TParam1, TParam2, TParam3, TFacade, TFacadeFactory>(
            Action<DiContainer, TParam1, TParam2, TParam3> facadeInstaller)
            where TFacade : IFacade
            where TFacadeFactory : FacadeFactory<TParam1, TParam2, TParam3, TFacade>
        {
            return this.Bind<TFacadeFactory>().ToMethod(
                x => x.Container.Instantiate<TFacadeFactory>(facadeInstaller));
        }

        ////////////// Other ////////////////

        // Walk the object graph for the given type
        // Should never throw an exception - returns them instead
        // Note: If you just want to know whether a binding exists for the given TContract,
        // use HasBinding instead
        public IEnumerable<ZenjectResolveException> ValidateResolve(InjectContext context)
        {
            ProviderBase provider = null;
            var result = TryGetUniqueProvider(context, out provider);

            if (result == DiContainer.ProviderLookupResult.Success)
            {
                Assert.IsNotNull(provider);

                if (ChecksForCircularDependencies)
                {
                    var lookupId = new LookupId(provider, context.BindingId);

                    // Allow one before giving up so that you can do circular dependencies via postinject or fields
                    if (_resolvesInProgress.Where(x => x.Equals(lookupId)).Count() > 1)
                    {
                        yield return new ZenjectResolveException(
                            "Circular dependency detected! \nObject graph:\n {0}".Fmt(context.GetObjectGraphString()));
                    }

                    _resolvesInProgress.Push(lookupId);
                    try
                    {
                        foreach (var error in provider.ValidateBinding(context))
                        {
                            yield return error;
                        }
                    }
                    finally
                    {
                        Assert.That(_resolvesInProgress.Peek().Equals(lookupId));
                        _resolvesInProgress.Pop();
                    }
                }
                else
                {
                    foreach (var error in provider.ValidateBinding(context))
                    {
                        yield return error;
                    }
                }
            }
            else if (result == DiContainer.ProviderLookupResult.Multiple)
            {
                yield return new ZenjectResolveException(
                    "Found multiple matches when only one was expected for dependency with type '{0}'{1} \nObject graph:\n{2}"
                    .Fmt(
                        context.MemberType.Name(),
                        (context.ObjectType == null ? "" : " when injecting into '{0}'".Fmt(context.ObjectType.Name())),
                        context.GetObjectGraphString()));
            }
            else
            {
                Assert.That(result == DiContainer.ProviderLookupResult.None);

                if (ReflectionUtil.IsGenericList(context.MemberType))
                {
                    var subType = context.MemberType.GetGenericArguments().Single();
                    var subContext = context.ChangeMemberType(subType);

                    var matches = GetAllProviderMatches(subContext);

                    if (matches.IsEmpty())
                    {
                        if (!context.Optional)
                        {
                            yield return new ZenjectResolveException(
                                "Could not find dependency with type 'List(0)'{1}.  If the empty list is also valid, you can allow this by using the [InjectOptional] attribute.' \nObject graph:\n{2}"
                                .Fmt(
                                    subContext.MemberType.Name(),
                                    (context.ObjectType == null ? "" : " when injecting into '{0}'".Fmt(context.ObjectType.Name())),
                                    context.GetObjectGraphString()));
                        }
                    }
                    else
                    {
                        foreach (var match in matches)
                        {
                            foreach (var error in match.ValidateBinding(context))
                            {
                                yield return error;
                            }
                        }
                    }
                }
                else
                {
                    if (!context.Optional)
                    {
                        yield return new ZenjectResolveException(
                            "Could not find required dependency with type '{0}'{1} \nObject graph:\n{2}"
                            .Fmt(
                                context.MemberType.Name(),
                                (context.ObjectType == null ? "" : " when injecting into '{0}'".Fmt(context.ObjectType.Name())),
                                context.GetObjectGraphString()));
                    }
                }
            }
        }

        public IEnumerable<ZenjectResolveException> ValidateObjectGraph<TConcrete>(params Type[] extras)
        {
            return ValidateObjectGraph(typeof(TConcrete), extras);
        }

        public IEnumerable<ZenjectResolveException> ValidateObjectGraph<TConcrete>(InjectContext context, params Type[] extras)
        {
            return ValidateObjectGraph(typeof(TConcrete), context, extras);
        }

        public IEnumerable<ZenjectResolveException> ValidateObjectGraph(
            Type contractType, params Type[] extras)
        {
            return ValidateObjectGraph(
                contractType, new InjectContext(this, contractType), extras);
        }

        public IEnumerable<ZenjectResolveException> ValidateObjectGraph(
            Type contractType, InjectContext context, params Type[] extras)
        {
            return ValidateObjectGraph(
                contractType, context, null, extras);
        }

        public IEnumerable<ZenjectResolveException> ValidateObjectGraph(
            Type concreteType, InjectContext currentContext, string concreteIdentifier, params Type[] extras)
        {
            if (concreteType.IsAbstract)
            {
                throw new ZenjectResolveException(
                    "Expected contract type '{0}' to be non-abstract".Fmt(concreteType.Name()));
            }

            var typeInfo = TypeAnalyzer.GetInfo(concreteType);
            var extrasList = extras.ToList();

            foreach (var dependInfo in typeInfo.AllInjectables)
            {
                Assert.IsEqual(dependInfo.ObjectType, concreteType);

                if (TryTakingFromExtras(dependInfo.MemberType, extrasList))
                {
                    continue;
                }

                var context = dependInfo.CreateInjectContext(this, currentContext, null, concreteIdentifier);

                foreach (var error in ValidateResolve(context))
                {
                    yield return error;
                }
            }

            if (!extrasList.IsEmpty())
            {
                yield return new ZenjectResolveException(
                    "Found unnecessary extra parameters passed when injecting into '{0}' with types '{1}'.  \nObject graph:\n{2}"
                    .Fmt(concreteType.Name(), String.Join(",", extrasList.Select(x => x.Name()).ToArray()), currentContext.GetObjectGraphString()));
            }
        }

        bool TryTakingFromExtras(Type contractType, List<Type> extrasList)
        {
            foreach (var extraType in extrasList)
            {
                if (extraType.DerivesFromOrEqual(contractType))
                {
                    var removed = extrasList.Remove(extraType);
                    Assert.That(removed);
                    return true;
                }
            }

            return false;
        }

        ////////////// Types ////////////////

        class ProviderPair
        {
            public readonly ProviderBase Provider;
            public readonly DiContainer Container;

            public ProviderPair(
                ProviderBase provider,
                DiContainer container)
            {
                Provider = provider;
                Container = container;
            }
        }

        public enum ProviderLookupResult
        {
            Success,
            Multiple,
            None
        }

        struct LookupId
        {
            public readonly ProviderBase Provider;
            public readonly BindingId BindingId;

            public LookupId(
                ProviderBase provider, BindingId bindingId)
            {
                Provider = provider;
                BindingId = bindingId;
            }
        }
    }
}
