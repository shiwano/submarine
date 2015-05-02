using UnityEngine;
using System;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class CommonInstaller : MonoInstaller
    {
        [Serializable]
        public class Settings
        {
            public GameObject ConnectionServicePrefab;
        }

        public Settings InstallerSettings;

        public override void InstallBindings()
        {
            Container.Bind<IInitializable>().ToSingle<CommonController>();
            Container.Bind<IDisposable>().ToSingle<CommonController>();
            Container.Bind<CommonController>().ToSingle();

            Container.Bind<ConnectionService>().ToSinglePrefab(InstallerSettings.ConnectionServicePrefab);
        }
    }
}
