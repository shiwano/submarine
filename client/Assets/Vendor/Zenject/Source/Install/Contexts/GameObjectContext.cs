#if !NOT_UNITY3D

using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;
using ModestTree.Util;
using UnityEngine;
using UnityEngine.Serialization;
using Zenject.Internal;

#pragma warning disable 649

namespace Zenject
{
    public class GameObjectContext : Context
    {
        readonly List<object> _dependencyRoots = new List<object>();

        [SerializeField]
        [Tooltip("Note that this field is optional and can be ignored in most cases.  This is really only needed if you want to control the 'Script Execution Order' of your subcontainer.  In this case, define a new class that derives from MonoKernel, add it to this game object, then drag it into this field.  Then you can set a value for 'Script Execution Order' for this new class and this will control when all ITickable/IInitializable classes bound within this subcontainer get called.")]
        [FormerlySerializedAs("_facade")]
        MonoKernel _kernel;

        DiContainer _container;

        public override DiContainer Container
        {
            get { return _container; }
        }

        public override IEnumerable<GameObject> GetRootGameObjects()
        {
            return new[] { this.gameObject };
        }

        [Inject]
        public void Construct(
            DiContainer parentContainer,
            [InjectOptional]
            InstallerExtraArgs installerExtraArgs)
        {
            Assert.IsNull(_container);

            _container = parentContainer.CreateSubContainer();

            foreach (var instance in GetInjectableMonoBehaviours().Cast<object>())
            {
                if (instance is MonoKernel)
                {
                    Assert.That(ReferenceEquals(instance, _kernel),
                        "Found MonoKernel derived class that is not hooked up to GameObjectContext.  If you use MonoKernel, you must indicate this to GameObjectContext by dragging and dropping it to the Kernel field in the inspector");
                }

                _container.QueueForInject(instance);
            }

            _container.IsInstalling = true;

            try
            {
                InstallBindings(installerExtraArgs);
            }
            finally
            {
                _container.IsInstalling = false;
            }

            Log.Debug("GameObjectContext: Injecting into child components...");

            _container.FlushInjectQueue();

            Assert.That(_dependencyRoots.IsEmpty());
            _dependencyRoots.AddRange(_container.ResolveDependencyRoots());

            if (_container.IsValidating)
            {
                // The root-level Container has its ValidateValidatables method
                // called explicitly - however, this is not so for sub-containers
                // so call it here instead
                _container.ValidateValidatables();
            }

            Log.Debug("GameObjectContext: Initialized successfully");
        }

        protected override IEnumerable<MonoBehaviour> GetInjectableMonoBehaviours()
        {
            // We inject on all components on the root except ourself
            foreach (var monoBehaviour in GetComponents<MonoBehaviour>())
            {
                if (monoBehaviour == null)
                {
                    // Missing script
                    continue;
                }

                if (monoBehaviour.GetType().DerivesFrom<MonoInstaller>())
                {
                    // Do not inject on installers since these are always injected before they are installed
                    continue;
                }

                if (monoBehaviour == this)
                {
                    continue;
                }

                yield return monoBehaviour;
            }

            foreach (var monoBehaviour in UnityUtil.GetDirectChildren(this.gameObject)
                .SelectMany<GameObject, MonoBehaviour>(ZenUtilInternal.GetInjectableMonoBehaviours))
            {
                yield return monoBehaviour;
            }
        }

        void InstallBindings(
            InstallerExtraArgs installerExtraArgs)
        {
            _container.DefaultParent = this.transform;

            _container.Bind<Context>().FromInstance(this);
            _container.Bind<GameObjectContext>().FromInstance(this);

            if (_kernel == null)
            {
                _container.Bind<MonoKernel>()
                    .To<DefaultGameObjectKernel>().FromNewComponentOn(this.gameObject).AsSingle().NonLazy();
            }
            else
            {
                _container.Bind<MonoKernel>().FromInstance(_kernel).AsSingle().NonLazy();
            }

            InstallSceneBindings();

            var extraArgsMap = new Dictionary<Type, List<TypeValuePair>>();

            if (installerExtraArgs != null)
            {
                extraArgsMap.Add(
                    installerExtraArgs.InstallerType, installerExtraArgs.ExtraArgs);
            }

            InstallInstallers();
        }

        public class InstallerExtraArgs
        {
            public Type InstallerType;
            public List<TypeValuePair> ExtraArgs;
        }
    }
}

#endif
