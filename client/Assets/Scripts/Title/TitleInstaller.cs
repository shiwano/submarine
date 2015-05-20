using UnityEngine;
using System;
using Zenject;
using UnityEngine.UI;

namespace Submarine
{
    public class TitleInstaller : MonoInstaller
    {
        [Serializable]
        public class Settings
        {
            public Button StartButton;
            public GameObject MatchingServicePrefab;
        }

        public Settings InstallerSettings;

        public override void InstallBindings()
        {
            Container.Bind<Settings>().ToSingleInstance(InstallerSettings);

            Container.Bind<IInitializable>().ToSingle<TitleController>();
            Container.Bind<IDisposable>().ToSingle<TitleController>();
            Container.Bind<TitleController>().ToSingle();

            Container.Bind<MatchingService>().ToSinglePrefab(InstallerSettings.MatchingServicePrefab);
        }
    }
}
