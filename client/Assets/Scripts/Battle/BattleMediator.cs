using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable, ITickable
    {
        [Inject]
        BattleService battleService;
        [Inject]
        BattleInputService battleInputService;
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
            battleService.Api.OnActorReceiveAsObservable().Subscribe(OnActorCreate).AddTo(view);

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

            battleInputService.IsAccelerating.Subscribe(OnAcceleratingChange).AddTo(view);
            battleInputService.OnTorpadeShootAsObservable().Subscribe(_ => OnTorpedoShoot()).AddTo(view);
            battleInputService.OnDecoyShootAsObservable().Subscribe(_ => OnDecoyShoot()).AddTo(view);
            battleInputService.OnLookoutShootAsObservable().Subscribe(_ => OnLookoutShoot()).AddTo(view);
            battleInputService.OnPingerUseAsObservable().Subscribe(_ => OnPingerUse()).AddTo(view);
        }

        void OnBattleFinish()
        {
            Logger.Log("Battle Finish");
            sceneChangeCommand.Execute(SceneNames.Lobby);
        }

        void OnAcceleratingChange(bool isAccelerating)
        {
            if (isAccelerating)
            {
                Logger.Log("Submarine accelerates");
                battleService.Api.SendAccelerationRequest();
            }
            else
            {
                Logger.Log("Submarine brakes");
                battleService.Api.SendBrakeRequest();
            }
        }

        void OnTorpedoShoot()
        {
            Logger.Log("Submarine shoots a torpedo");
            battleService.Api.SendTorpedoRequest();
        }

        void OnDecoyShoot()
        {
            Logger.Log("Sumarine shoots a decoy");
        }

        void OnLookoutShoot()
        {
            Logger.Log("Submarine shoots a lookout");
        }

        void OnPingerUse()
        {
            Logger.Log("Submarine uses pinger");
            battleService.Api.SendPingerRequest();
        }

        void OnActorCreate(Type.Battle.Actor actor)
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
