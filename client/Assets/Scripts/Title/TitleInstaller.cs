using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine
{
    public class TitleInstaller : MonoInstaller
    {
        [SerializeField]
        TitleView view;

        public override void InstallBindings()
        {
            Container.Bind<Services.Authentication>().ToSingle();

            Container.Bind<Events.Title>().ToSingle();
            Container.BindCommand<Commands.Login>().HandleWithSingle<Commands.Login.Handler>();
            Container.BindCommand<Commands.SignUp, string>().HandleWithSingle<Commands.SignUp.Handler>();

            Container.Bind<TitleView>().ToSingleInstance(view);
            Container.Bind<TitleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<TitleMediator>();
        }
    }
}
