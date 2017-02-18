using UnityEngine;
using Zenject;

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
            Container.Bind<TitleEvent.SignUpStart>().AsSingle();

            Container.Bind<AuthenticationService>().AsSingle();

            Container.DeclareSignal<LoginCommand>().RequireHandler();
            Container.BindSignal<LoginCommand>().To<LoginCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<SignUpCommand>().RequireHandler();
            Container.BindSignal<string, SignUpCommand>().To<SignUpCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<DeleteLoginDataCommand>().RequireHandler();
            Container.BindSignal<DeleteLoginDataCommand>().To<DeleteLoginDataCommand.Handler>(x => x.Execute).AsSingle();

            Container.BindInstance(titleView);
            Container.BindInterfacesAndSelfTo<TitleMediator>().AsSingle();

            Container.BindInstance(signUpView);
            Container.BindInterfacesAndSelfTo<SignUpMediator>().AsSingle();
        }
    }
}
