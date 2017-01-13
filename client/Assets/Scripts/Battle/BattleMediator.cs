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
        BattleModel battleModel;
        [Inject]
        BattleView view;
        [Inject]
        InitializeBattleCommand initializeBattleCommand;
        [Inject]
        ActorContainer actorContainer;

        public void Initialize()
        {
            battleModel.OnPrepareAsObservable().Subscribe(_ => OnBattlePrepare()).AddTo(view);
            battleModel.OnStartAsObservable().Subscribe(_ => OnBattleStart()).AddTo(view);
            battleModel.OnFinishAsObservable().Subscribe(_ => OnBattleFinish()).AddTo(view);
            battleModel.ActorsById.ObserveAdd().Subscribe(e => OnActorAdd(e.Value)).AddTo(view);
            battleModel.ActorsById.ObserveRemove().Subscribe(x => OnActorRemove(x.Value)).AddTo(view);

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
        }

        void OnBattleStart()
        {
            Logger.Log("Battle Start");
        }

        void OnBattleFinish()
        {
            Logger.Log("Battle Finish");
        }

        void OnActorAdd(Type.Battle.Actor actor)
        {
            var actorFacade = actorContainer.CreateActor(actor);
            actorCreateEvent.Invoke(actorFacade);
        }

        void OnActorRemove(Type.Battle.Actor actor)
        {
            ActorFacade actorFacade;
            if (actorContainer.TryDestroyActor(actor.Id, out actorFacade))
            {
                actorDestroyEvent.Invoke(actorFacade);
            }
        }
    }
}
