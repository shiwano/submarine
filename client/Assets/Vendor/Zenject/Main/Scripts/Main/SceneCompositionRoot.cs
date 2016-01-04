#if !ZEN_NOT_UNITY3D

#pragma warning disable 414
using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;
using ModestTree.Util;
using UnityEngine;

#if UNITY_5_3
using UnityEngine.SceneManagement;
#endif

namespace Zenject
{
    public sealed class SceneCompositionRoot : CompositionRoot
    {
        public static Action<DiContainer> BeforeInstallHooks;
        public static Action<DiContainer> AfterInstallHooks;

        public bool OnlyInjectWhenActive = true;

        [SerializeField]
        public MonoInstaller[] Installers = new MonoInstaller[0];

        DiContainer _container;
        IFacade _rootFacade = null;

        static List<IInstaller> _staticInstallers = new List<IInstaller>();

        public override DiContainer Container
        {
            get
            {
                return _container;
            }
        }

        public override IFacade RootFacade
        {
            get
            {
                return _rootFacade;
            }
        }

        string GetCurrentSceneName()
        {
#if UNITY_5_3
            return SceneManager.GetActiveScene().name;
#else
            return Application.loadedLevelName;
#endif
        }

        protected override void Initialize()
        {
            Log.Debug("Initializing SceneCompositionRoot in scene '{0}'", GetCurrentSceneName());
            InitContainer();
            Log.Debug("SceneCompositionRoot: Finished install phase.  Injecting into scene...");
            InitialInject();

            Log.Debug("SceneCompositionRoot: Resolving root IFacade...");
            _rootFacade = _container.Resolve<IFacade>();
        }

        public void Start()
        {
            // Always run the IInitializable's at the very beginning of Start()
            // This file (SceneCompositionRoot) should always have the earliest execution order (see SceneCompositionRoot.cs.meta)
            // This is necessary in some edge cases where parts of Unity do not work the same during Awake() as they do in Start/Update
            // For example, changing rectTransform.localPosition does not automatically update rectTransform.position in some cases
            // Also, most people treat Awake() as very minimal initialization, such as setting up a valid state for the
            // object, initializing variables, etc. and treat Start() as the place where more complex initialization occurs,
            // so this is consistent with that convention as well
            GlobalCompositionRoot.Instance.InitializeRootIfNecessary();
            _rootFacade.Initialize();
        }

        // This method is used for cases where you need to create the SceneCompositionRoot entirely in code
        // Necessary because the Awake() method is called immediately after AddComponent<SceneCompositionRoot>
        // so there's no other way to add installers to it
        public static SceneCompositionRoot AddComponent(
            GameObject gameObject, IInstaller rootInstaller)
        {
            return AddComponent(gameObject, new List<IInstaller>() { rootInstaller });
        }

        public static SceneCompositionRoot AddComponent(
            GameObject gameObject, List<IInstaller> installers)
        {
            Assert.That(_staticInstallers.IsEmpty());
            _staticInstallers.AddRange(installers);
            return gameObject.AddComponent<SceneCompositionRoot>();
        }

        void InitContainer()
        {
            _container = CreateContainer(
                false, GlobalCompositionRoot.Instance.Container, _staticInstallers);
            _staticInstallers.Clear();
        }

        void InitialInject()
        {
            var rootGameObjects = GameObject.FindObjectsOfType<Transform>()
                .Where(x => x.parent == null && x.GetComponent<GlobalCompositionRoot>() == null && (x.GetComponent<SceneCompositionRoot>() == null || x == this.transform))
                .Select(x => x.gameObject).ToList();

            foreach (var rootObj in rootGameObjects)
            {
                _container.InjectGameObject(rootObj, true, !OnlyInjectWhenActive);
            }
        }

        public DiContainer CreateContainer(
            bool isValidating, DiContainer parentContainer, List<IInstaller> extraInstallers)
        {
            var container = new DiContainer(this.transform, parentContainer);

            container.IsValidating = isValidating;

            container.Bind<SceneCompositionRoot>().ToInstance(this);
            container.Bind<CompositionRoot>().ToInstance(this);

            if (BeforeInstallHooks != null)
            {
                BeforeInstallHooks(container);
                // Reset extra bindings for next time we change scenes
                BeforeInstallHooks = null;
            }

            container.Install<StandardInstaller>();

            var allInstallers = extraInstallers.Concat(Installers).ToList();

            if (allInstallers.Where(x => x != null).IsEmpty())
            {
                Log.Warn("No installers found while initializing SceneCompositionRoot");
            }
            else
            {
                container.Install(allInstallers);
            }

            if (AfterInstallHooks != null)
            {
                AfterInstallHooks(container);
                // Reset extra bindings for next time we change scenes
                AfterInstallHooks = null;
            }

            return container;
        }
    }
}

#endif
