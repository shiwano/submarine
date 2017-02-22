using UniRx;
using Zenject;

namespace Submarine.Battle
{
    public class RadarMediator : MediatorBase<RadarView>, IInitializable
    {
        [Inject]
        BattleEvent.ActorCreate actorCreateEvent;
        [Inject]
        BattleEvent.ActorDestroy actorDestroyEvent;

        public void Initialize()
        {
            actorCreateEvent.AsObservable().Subscribe(OnActorCreate).AddTo(view);
            actorDestroyEvent.AsObservable().Subscribe(OnActorDestroy).AddTo(view);
        }

        void OnActorCreate(ActorFacade actorFacade)
        {
            view.CreatePin(actorFacade);
        }

        void OnActorDestroy(ActorFacade actorFacade)
        {
            view.DestroyPin(actorFacade);
        }
    }
}
