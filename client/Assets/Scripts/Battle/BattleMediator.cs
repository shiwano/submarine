using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable, ITickable
    {
        [Inject]
        BattleModel battleModel;
        [Inject]
        BattleView view;
        [Inject]
        InitializeBattleCommand initializeBattleCommand;
        [Inject]
        SceneChangeCommand sceneChangeCommand;
        [Inject]
        ActorContainer actorContainer;
        [Inject]
        ThirdPersonCamera thirdPersonCamera;

        public void Initialize()
        {
            battleModel.OnPrepareAsObservable().Take(1).Subscribe(_ => OnBattlePrepare()).AddTo(view);
            battleModel.OnStartAsObservable().Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            battleModel.OnFinishAsObservable().Take(1).Subscribe(_ => OnBattleFinish()).AddTo(view);
            battleModel.Actors.ObserveAdd().Subscribe(e => OnActorCreated(e.Value)).AddTo(view);

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
            sceneChangeCommand.Execute(SceneNames.Lobby);
        }

        void OnActorCreated(Type.Battle.Actor actor)
        {
            switch (actor.Type)
            {
                case Type.Battle.ActorType.Submarine:
                {
                    var submarine = actorContainer.CreateSubmarine(actor);
                    if (submarine.IsMine)
                    {
                        thirdPersonCamera.SetTarget(submarine.View.transform);
                    }
                    break;
                }
            }
        }
    }
}
