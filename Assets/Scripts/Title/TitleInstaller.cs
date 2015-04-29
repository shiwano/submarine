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
        }

        public Settings SceneSettings;

        public override void InstallBindings()
        {
            Container.Bind<Settings>().ToSingleInstance(SceneSettings);

            Container.Bind<IInitializable>().ToSingle<TitleController>();
            Container.Bind<TitleController>().ToSingle();
        }
    }
}
