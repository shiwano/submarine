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

            Container.Bind<Models.User>().ToSingle();
            Container.Bind<Services.PermanentDataStore>().ToSingle();

            Container.BindCommand<Commands.SceneChange, SceneNames>().HandleWithSingle<Commands.SceneChange.Handler>();
            Container.BindCommand<Commands.ApplicationStart>().HandleWithSingle<Commands.ApplicationStart.Handler>();
            Container.BindCommand<Commands.ApplicationPause>().HandleWithSingle<Commands.ApplicationPause.Handler>();
            Container.BindCommand<Commands.ApplicationQuit>().HandleWithSingle<Commands.ApplicationQuit.Handler>();

            Container.Bind<IInitializable>().ToSingle<CommonMediator>();
            Container.Bind<CommonMediator>().ToSingle();
        }
    }
}
