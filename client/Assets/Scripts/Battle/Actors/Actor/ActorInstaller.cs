using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public abstract class ActorInstaller<TFacade, TView> : MonoInstaller
        where TFacade : ActorFacade
        where TView : ActorView
    {
        [Inject]
        Type.Battle.Actor actor;

        public override void InstallBindings()
        {
            Container.BindInstance(actor);
            Container.Bind<ActorMotor>().AsSingle();

            Container.BindInterfacesAndSelfTo<TFacade>().AsSingle();

            var view = GetComponent<TView>();
            Container.Bind<TView>().FromInstance(view).AsSingle();
            Container.Bind<ActorView>().FromInstance(view).AsSingle();
        }
    }
}
