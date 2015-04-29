using UnityEngine;
using System;
using Zenject;
using UnityEngine.UI;

namespace Submarine
{
    public class BattleInstaller : MonoInstaller
    {
        [Serializable]
        public class Settings
        {
            public GameObject BattleServicePrefab;
        }

        public Settings InstallerSettings;

        public override void InstallBindings()
        {
            Container.Bind<Settings>().ToSingleInstance(InstallerSettings);

            Container.Bind<BattleService>().ToSinglePrefab(InstallerSettings.BattleServicePrefab);

            Container.Bind<IInitializable>().ToSingle<TitleController>();
            Container.Bind<BattleController>().ToSingle();
        }
    }
}
