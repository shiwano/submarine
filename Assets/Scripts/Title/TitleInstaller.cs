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
            public SubmarineSettings Submarine;
            public UISettings UI;

            [Serializable]
            public class SubmarineSettings
            {
                public GameObject Prefab;
                public Vector3 StartPosition;
                public Vector3 StartRotation;
            }

            [Serializable]
            public class UISettings
            {
                public Button StartButton;
            }
        }

        public Settings SceneSettings;

        public override void InstallBindings()
        {
            Container.Bind<Settings>().ToSingleInstance(SceneSettings);

            Container.Bind<IInitializable>().ToSingle<TitleController>();
            Container.Bind<TitleController>().ToSingle();

            Container.Bind<Transform>("Submarine").ToSinglePrefab(SceneSettings.Submarine.Prefab);
        }
    }
}
