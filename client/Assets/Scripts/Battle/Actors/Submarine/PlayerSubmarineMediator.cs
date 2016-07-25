using UnityEngine;
using System;
using UniRx;
using Zenject;

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
        SubmarineView submarineView;

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

            battleInputService.OnTorpadeShootAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnTorpedoShoot())
                .AddTo(disposables);

            battleInputService.OnDecoyShootAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnDecoyShoot())
                .AddTo(disposables);

            battleInputService.OnLookoutShootAsObservable()
                .Where(_ => battleModel.IsInBattle)
                .Subscribe(_ => OnLookoutShoot())
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
    }
}
