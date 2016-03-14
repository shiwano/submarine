using UnityEngine;
using System;
using UniRx;
using Zenject;

namespace Submarine.Battle
{
    public class PlayerSubmarineMediator : IInitializable, ITickable, IDisposable
    {
        [Inject]
        BattleEvent.PlayerSubmarineCreate playerSubmarineCreateEvent;
        [Inject]
        BattleModel battleModel;
        [Inject]
        BattleService battleService;
        [Inject]
        BattleInputService battleInputService;
        [Inject]
        ThirdPersonCamera thirdPersonCamera;

        readonly CompositeDisposable disposables = new CompositeDisposable();
        SubmarineFacade submarine;

        readonly float minSignificantDirection = 0.05f;
        readonly float directionThresholdForSendingDirection = 10f;
        readonly TimeSpan durationThresholdForSendingDirection = TimeSpan.FromSeconds(1.5f);
        double lastSentDirection;
        DateTime lastSentDirectionAt;

        public void Initialize()
        {
            playerSubmarineCreateEvent.AsObservable().Subscribe(OnPlayerSubmarineCreate).AddTo(disposables);
        }

        public void Tick()
        {
            if (submarine != null)
            {
                submarine.Turn(battleInputService.TurningRate);
                SendDirectionIfNeeded(submarine.Direction);
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

        void OnPlayerSubmarineCreate(SubmarineFacade submarine)
        {
            this.submarine = submarine;
            thirdPersonCamera.SetTarget(submarine.View.transform);

            battleInputService.IsAccelerating.Subscribe(OnAcceleratingChange).AddTo(disposables);
            battleInputService.OnTorpadeShootAsObservable().Subscribe(_ => OnTorpedoShoot()).AddTo(disposables);
            battleInputService.OnDecoyShootAsObservable().Subscribe(_ => OnDecoyShoot()).AddTo(disposables);
            battleInputService.OnLookoutShootAsObservable().Subscribe(_ => OnLookoutShoot()).AddTo(disposables);
            battleInputService.OnPingerUseAsObservable().Subscribe(_ => OnPingerUse()).AddTo(disposables);
        }

        void OnAcceleratingChange(bool isAccelerating)
        {
            var direction = submarine.Direction;

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
