using Zenject;
using Zenject.Commands;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class CommonInstaller : MonoInstaller
    {
        public override void InstallBindings()
        {
            var config = Type.Config.Load();
            var webApi = new TyphenApi.WebApi.Submarine(config.WebApiServerBaseUri);

            Container.Bind<Type.Config>().ToSingleInstance(config);
            Container.Bind<TyphenApi.WebApi.Submarine>().ToSingleInstance(webApi);

            Container.Bind<UserModel>().ToSingle();
            Container.Bind<PermanentDataStoreService>().ToSingle();

            Container.BindCommand<SceneChangeCommand, SceneNames>().HandleWithSingle<SceneChangeCommand.Handler>();
            Container.BindCommand<ApplicationStartCommand>().HandleWithSingle<ApplicationStartCommand.Handler>();
            Container.BindCommand<ApplicationPauseCommand>().HandleWithSingle<ApplicationPauseCommand.Handler>();
            Container.BindCommand<ApplicationQuitCommand>().HandleWithSingle<ApplicationQuitCommand.Handler>();

            Container.Bind<IInitializable>().ToSingle<CommonMediator>();
            Container.Bind<CommonMediator>().ToSingle();
        }
    }
}
