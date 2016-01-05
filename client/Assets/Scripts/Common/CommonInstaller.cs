using UnityEngine;
using System;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class CommonInstaller : MonoInstaller
    {
        public override void InstallBindings()
        {
            Container.Bind<Services.LoginApi>().ToSingle();
            Container.Bind<Services.PermanentDataStore>().ToSingle();

            Container.Bind<Commands.SceneChange>().ToSingle();
            Container.Bind<Commands.ApplicationStart>().ToSingle().WhenInjectedInto<CommonMediator>();
            Container.Bind<Commands.ApplicationPause>().ToSingle().WhenInjectedInto<CommonMediator>();
            Container.Bind<Commands.ApplicationQuit>().ToSingle().WhenInjectedInto<CommonMediator>();

            Container.Bind<IInitializable>().ToSingle<CommonMediator>();
            Container.Bind<CommonMediator>().ToSingle();
        }
    }
}
