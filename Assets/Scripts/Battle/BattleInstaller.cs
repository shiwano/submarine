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
            public BattleService BattleService;
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

            Container.Bind<BattleService>().ToSingleInstance(InstallerSettings.BattleService);

            Container.Bind<IInitializable>().ToSingle<BattleInput>();
            Container.Bind<IDisposable>().ToSingle<BattleInput>();
            Container.Bind<BattleInput>().ToSingle();

            Container.Bind<IInitializable>().ToSingle<BattleController>();
            Container.Bind<IDisposable>().ToSingle<BattleController>();
            Container.Bind<BattleController>().ToSingle();

            Container.Bind<IInitializable>().ToSingle<BattleObjectSpawner>();
            Container.Bind<IDisposable>().ToSingle<BattleObjectSpawner>();
            Container.Bind<ITickable>().ToSingle<BattleObjectSpawner>();
            Container.Bind<BattleObjectSpawner>().ToSingle();

            Container.Bind<ITickable>().ToSingle<ThirdPersonCamera>();
            Container.Bind<ThirdPersonCamera>().ToSingle<ThirdPersonCamera>();

            Container.Bind<SubmarineFactory>().ToSingle();
            Container.Bind<TorpedoFactory>().ToSingle();
        }
    }
}
