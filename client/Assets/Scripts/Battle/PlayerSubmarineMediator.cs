using System;
using UniRx;
using Zenject;

namespace Submarine.Battle
{
    public class PlayerSubmarineMediator : IInitializable, ITickable, IDisposable
    {
        [Inject]
        BattleEvents.PlayerSubmarineCreate playerSubmarineCreateEvent;
        [Inject]
        BattleService battleService;
        [Inject]
        BattleInputService battleInputService;
        [Inject]
        ThirdPersonCamera thirdPersonCamera;

        readonly CompositeDisposable disposables = new CompositeDisposable();
        SubmarineFacade submarine;

        public void Initialize()
        {
            playerSubmarineCreateEvent.AsObservable().Subscribe(OnPlayerSubmarineCreate).AddTo(disposables);
        }

        public void Tick()
        {
            if (submarine != null)
            {
                submarine.Turn(battleInputService.TurningRate);
            }
        }

        public void Dispose()
        {
            disposables.Dispose();
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
    }
}
