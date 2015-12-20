#if !ZEN_NOT_UNITY3D

#pragma warning disable 414
using ModestTree;

using System;
using System.Collections.Generic;
using System.Linq;
using UnityEngine;

namespace Zenject
{
    public sealed class GlobalCompositionRoot : CompositionRoot
    {
        static GlobalCompositionRoot _instance;
        DiContainer _container;
        IFacade _rootFacade;
        bool _hasInitialized;

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

        public static GlobalCompositionRoot Instance
        {
            get
            {
                if (_instance == null)
                {
                    _instance = new GameObject("Global Composition Root")
                        .AddComponent<GlobalCompositionRoot>();
                }
                return _instance;
            }
        }

        protected override void Initialize()
        {
            DontDestroyOnLoad(gameObject);

            // Is this a good idea?
            //go.hideFlags = HideFlags.HideInHierarchy;

            _container = CreateContainer(false, this);
            _rootFacade = _container.Resolve<IFacade>();
        }

        public void InitializeRootIfNecessary()
        {
            if (!_hasInitialized)
            {
                _hasInitialized = true;
                _rootFacade.Initialize();
            }
        }

        public static DiContainer CreateContainer(bool allowNullBindings, GlobalCompositionRoot root)
        {
            Assert.That(allowNullBindings || root != null);

            var container = new DiContainer(root == null ? null : root.transform);

            container.AllowNullBindings = allowNullBindings;

            container.Bind<GlobalCompositionRoot>().ToInstance(root);
            container.Bind<CompositionRoot>().ToInstance(root);

            container.Install<StandardInstaller>();

            container.Install(GetGlobalInstallers());

            return container;
        }

        static IEnumerable<IInstaller> GetGlobalInstallers()
        {
            // Allow either naming convention
            var installerConfigs1 = Resources.LoadAll("ZenjectGlobalCompositionRoot", typeof(GlobalInstallerConfig));
            var installerConfigs2 = Resources.LoadAll("ZenjectGlobalInstallers", typeof(GlobalInstallerConfig));

            return installerConfigs1.Concat(installerConfigs2).Cast<GlobalInstallerConfig>().SelectMany(x => x.Installers).Cast<IInstaller>();
        }
    }
}

#endif
