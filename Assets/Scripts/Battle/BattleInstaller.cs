using UnityEngine;
using System;
using System.Collections.Generic;
using Zenject;
using UnityEngine.UI;

namespace Submarine
{
    public class BattleInstaller : MonoInstaller
    {
        [Serializable]
        public class Settings
        {
            public Camera MainCamera;
            public GameObject BattleServicePrefab;
            public SubmarineSettings Submarine;

            [Serializable]
            public class SubmarineSettings
            {
                public List<Vector3> StartPositions = new List<Vector3>();
            }
        }

        public Settings InstallerSettings;

        public override void InstallBindings()
        {
            Container.Bind<Settings>().ToSingleInstance(InstallerSettings);
            Container.Bind<Camera>("MainCamera").ToSingleInstance(InstallerSettings.MainCamera);

            Container.Bind<BattleService>().ToSinglePrefab(InstallerSettings.BattleServicePrefab);
            Container.Bind<BattleInput>().ToSingle();

            Container.Bind<IInitializable>().ToSingle<BattleController>();
            Container.Bind<IDisposable>().ToSingle<BattleController>();
            Container.Bind<BattleController>().ToSingle();

            Container.Bind<IInitializable>().ToSingle<ThirdPersonCamera>();
            Container.Bind<ITickable>().ToSingle<ThirdPersonCamera>();
            Container.Bind<ThirdPersonCamera>().ToSingle<ThirdPersonCamera>();

            Container.Bind<SubmarineFactory>().ToSingle();
        }
    }
}
