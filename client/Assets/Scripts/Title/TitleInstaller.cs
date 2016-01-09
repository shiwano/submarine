using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine.Title
{
    public class TitleInstaller : MonoInstaller
    {
        [SerializeField]
        TitleView view;

        public override void InstallBindings()
        {
            Container.Bind<AuthenticationService>().ToSingle();

            Container.Bind<TitleEvents>().ToSingle();
            Container.BindCommand<LoginCommand>().HandleWithSingle<LoginCommand.Handler>();
            Container.BindCommand<SignUpCommand, string>().HandleWithSingle<SignUpCommand.Handler>();
            Container.BindCommand<DeleteLoginDataCommand>().HandleWithSingle<DeleteLoginDataCommand.Handler>();

            Container.Bind<TitleView>().ToSingleInstance(view);
            Container.Bind<TitleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<TitleMediator>();
        }
    }
}
