using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable, ITickable
    {
        [Inject]
        BattleEvent.ActorCreate actorCreateEvent;
        [Inject]
        BattleEvent.ActorDestroy actorDestroyEvent;
        [Inject]
        BattleService battleService;
        [Inject]
        BattleModel battleModel;
        [Inject]
        BattleView view;
        [Inject]
        InitializeBattleCommand initializeBattleCommand;
        [Inject]
        ActorContainer actorContainer;

        public void Initialize()
        {
            battleModel.OnPrepareAsObservable().Take(1).Subscribe(_ => OnBattlePrepare()).AddTo(view);
            battleModel.OnStartAsObservable().Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            battleModel.OnFinishAsObservable().Take(1).Subscribe(_ => OnBattleFinish()).AddTo(view);
            initializeBattleCommand.Execute();
        }

        public void Tick()
        {
            if (battleModel.IsInBattle)
            {
                view.ElapsedTime = battleModel.ElapsedTime;
            }
        }

        void OnBattlePrepare()
        {
            Logger.Log("Battle Prepare");
            battleService.Api.OnActorReceiveAsObservable().Subscribe(OnActorCreate).AddTo(view);
            battleService.Api.OnMovementReceiveAsObservable().Subscribe(OnActorMove).AddTo(view);
            battleService.Api.OnDestructionReceiveAsObservable().Subscribe(OnActorDestroy).AddTo(view);
        }

        void OnBattleStart()
        {
            Logger.Log("Battle Start");
        }

        void OnBattleFinish()
        {
            Logger.Log("Battle Finish");
        }

        void OnActorCreate(Type.Battle.Actor actor)
        {
            var actorFacade = actorContainer.CreateActor(actor);
            actorCreateEvent.Invoke(actorFacade);
        }

        void OnActorMove(Type.Battle.Movement movement)
        {
            var actor = actorContainer.Get(movement.ActorId);
            if (actor != null)
            {
                actor.Motor.SetMovement(movement);
            }
        }

        void OnActorDestroy(Type.Battle.Destruction destruction)
        {
            ActorFacade actorFacade;
            if (actorContainer.TryDestroyActor(destruction.ActorId, out actorFacade))
            {
                actorDestroyEvent.Invoke(actorFacade);
            }
        }
    }
}
