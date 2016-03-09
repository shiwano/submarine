using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine.Title
{
    public class TitleInstaller : MonoInstaller
    {
        [SerializeField]
        TitleView titleView;
        [SerializeField]
        SignUpView signUpView;

        public override void InstallBindings()
        {
            Container.Bind<TitleEvent.SignUpStart>().ToSingle();

            Container.Bind<AuthenticationService>().ToSingle();

            Container.BindCommand<LoginCommand>().HandleWithSingle<LoginCommand.Handler>();
            Container.BindCommand<SignUpCommand, string>().HandleWithSingle<SignUpCommand.Handler>();
            Container.BindCommand<DeleteLoginDataCommand>().HandleWithSingle<DeleteLoginDataCommand.Handler>();

            Container.Bind<TitleView>().ToSingleInstance(titleView);
            Container.Bind<TitleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<TitleMediator>();

            Container.Bind<SignUpView>().ToSingleInstance(signUpView);
            Container.Bind<SignUpMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<SignUpMediator>();
        }
    }
}
