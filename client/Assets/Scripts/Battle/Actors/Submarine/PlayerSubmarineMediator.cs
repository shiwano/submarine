using System;
using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class PlayerSubmarineMediator : IInitializable, ITickable, IDisposable
    {
        [Inject]
        BattleModel battleModel;
        [Inject]
        BattleService battleService;
        [Inject]
        BattleInputService battleInputService;
        [Inject]
        ThirdPersonCamera thirdPersonCamera;
        [Inject]
        Type.Battle.Actor actor;
        [Inject]
        SubmarineView submarineView;
        [Inject]
        EquipmentView equipmentView;

        readonly CompositeDisposable disposables = new CompositeDisposable();

        readonly float minSignificantDirection = 0.05f;
        readonly float directionThresholdForSendingDirection = 10f;
        readonly TimeSpan durationThresholdForSendingDirection = TimeSpan.FromSeconds(1.5f);
        double lastSentDirection;
        DateTime lastSentDirectionAt;

        public void Initialize()
        {
            thirdPersonCamera.SetTarget(submarineView.transform);

            battleInputService.IsAccelerating
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(OnAcceleratingChange)
                .AddTo(disposables);

            battleInputService.OnTorpadeUseAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnTorpedoUse())
                .AddTo(disposables);

            battleInputService.OnDecoyUseAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnDecoyUse())
                .AddTo(disposables);

            battleInputService.OnWatcherUseAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnWatcherUse())
                .AddTo(disposables);

            battleInputService.OnPingerUseAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnPingerUse())
                .AddTo(disposables);
        }

        public void Tick()
        {
            if (battleModel.IsInBattle)
            {
                equipmentView.Refresh(battleModel.Now, actor.Submarine.Equipment);
                submarineView.Turn(battleInputService.TurningRate);
                SendDirectionIfNeeded(submarineView.ActorDirection);
            }
        }

        public void Dispose()
        {
            disposables.Dispose();
        }

        void SendDirectionIfNeeded(double direction)
        {
            var diff = Math.Abs(direction - lastSentDirection);

            if (diff > directionThresholdForSendingDirection ||
                (diff > minSignificantDirection &&
                 battleModel.Now - lastSentDirectionAt > durationThresholdForSendingDirection))
            {
                battleService.Api.SendTurnRequest(direction);
                lastSentDirection = direction;
                lastSentDirectionAt = battleModel.Now;
            }
        }

        void OnAcceleratingChange(bool isAccelerating)
        {
            var direction = submarineView.ActorDirection;

            if (isAccelerating)
            {
                Logger.Log("Submarine accelerates");
                battleService.Api.SendAccelerationRequest(direction);
                lastSentDirection = direction;
                lastSentDirectionAt = battleModel.Now;
            }
            else
            {
                Logger.Log("Submarine brakes");
                battleService.Api.SendBrakeRequest(direction);
                lastSentDirection = direction;
                lastSentDirectionAt = battleModel.Now;
            }
        }

        void OnTorpedoUse()
        {
            Logger.Log("Submarine shoots a torpedo");
            battleService.Api.SendTorpedoRequest();
        }

        void OnDecoyUse()
        {
            Logger.Log("Sumarine shoots a decoy");
        }

        void OnWatcherUse()
        {
            Logger.Log("Submarine uses a watcher");
            battleService.Api.SendWatcherRequest();
        }

        void OnPingerUse()
        {
            Logger.Log("Submarine uses pinger");
            battleService.Api.SendPingerRequest();
        }
    }
}
