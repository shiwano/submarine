using UnityEngine;
using System;
using Zenject;
using Zenject.Commands;
using UnityEngine.UI;

namespace Submarine
{
    public class TitleInstaller : MonoInstaller
    {
        [SerializeField]
        TitleView view;

        public override void InstallBindings()
        {
            Container.Bind<Services.LoginApi>().ToSingle();

            Container.Bind<TitleEvent>().ToSingle();
            Container.BindCommand<Commands.Login>().HandleWithSingle<Commands.Login.Handler>();

            Container.Bind<TitleView>().ToSingleInstance(view);
            Container.Bind<TitleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<TitleMediator>();
        }
    }
}
