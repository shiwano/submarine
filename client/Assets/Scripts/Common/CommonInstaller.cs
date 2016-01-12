using Zenject;
using Zenject.Commands;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class CommonInstaller : MonoInstaller
    {
        public override void InstallBindings()
        {
            Container.Bind<Type.Config>().ToSingleInstance(Type.Config.Load());
            Container.Bind<TyphenApi.WebApi.Submarine>().ToSingle();

            Container.Bind<UserModel>().ToSingle();
            Container.Bind<LobbyModel>().ToSingle();
            Container.Bind<PermanentDataStoreService>().ToSingle();
            Container.Bind<BattleService>().ToSingle();

            Container.BindCommand<SceneChangeCommand, SceneNames>().HandleWithSingle<SceneChangeCommand.Handler>();
            Container.BindCommand<ApplicationStartCommand>().HandleWithSingle<ApplicationStartCommand.Handler>();
            Container.BindCommand<ApplicationPauseCommand>().HandleWithSingle<ApplicationPauseCommand.Handler>();
            Container.BindCommand<ApplicationQuitCommand>().HandleWithSingle<ApplicationQuitCommand.Handler>();

            Container.Bind<CommonMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<CommonMediator>();
        }
    }
}
