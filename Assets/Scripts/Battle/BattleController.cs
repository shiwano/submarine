using UnityEngine;
using System;
using System.Linq;
using UniRx;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable, ITickable
    {
        readonly BattleInstaller.Settings settings;
        readonly ConnectionService connection;
        readonly BattleService battleService;
        readonly BattleObjectContainer objectContainer;
        readonly ThirdPersonCamera thirdPersonCamera;
        readonly Radar radar;
        readonly CompositeDisposable disposables = new CompositeDisposable();

        PlayerSubmarine playerSubmarine;

        public BattleController(
            BattleInstaller.Settings settings,
            ConnectionService connection,
            BattleService battleService,
            BattleObjectContainer objectContainer,
            ThirdPersonCamera thirdPersonCamera,
            Radar radar)
        {
            this.settings = settings;
            this.connection = connection;
            this.battleService = battleService;
            this.objectContainer = objectContainer;
            this.thirdPersonCamera = thirdPersonCamera;
            this.radar = radar;
        }

        public void Initialize()
        {
            if (!connection.InRoom) { return; }

            battleService.ResultDecided += OnResultDecided;
            settings.UI.Victory.OnClickAsObservable()
                .Subscribe(_ => MoveToTitle())
                .AddTo(disposables);
            settings.UI.Defeat.OnClickAsObservable()
                .Subscribe(_ => MoveToTitle())
                .AddTo(disposables);

            battleService.StartBattle();

            playerSubmarine = objectContainer.SpawnPlayerSubmarine(settings.Map.StartPositions[connection.Player.ID]);
            thirdPersonCamera.SetTarget(playerSubmarine.Hooks.transform);

            playerSubmarine.Resources.Decoy.CanUse
                .SubscribeToInteractable(settings.UI.DecoyButton)
                .AddTo(disposables);
            playerSubmarine.Resources.Decoy.CountDown
                .Select(i => i <= 0 ? "" : i.ToString())
                .Subscribe(s => settings.UI.DecoyCoolDown.text = s)
                .AddTo(disposables);
            playerSubmarine.Resources.Pinger.CanUse
                .SubscribeToInteractable(settings.UI.PingerButton)
                .AddTo(disposables);
            playerSubmarine.Resources.Pinger.CountDown
                .Select(i => i <= 0 ? "" : i.ToString())
                .Subscribe(s => settings.UI.PingerCoolDown.text = s)
                .AddTo(disposables);
            playerSubmarine.Resources.Pinger.IsUsing
                .Subscribe(radar.SetPinger)
                .AddTo(disposables);
            playerSubmarine.Resources.Lookout.CanUse
                .SubscribeToInteractable(settings.UI.LookoutButton)
                .AddTo(disposables);
            playerSubmarine.Resources.Lookout.CountDown
                .Select(i => i <= 0 ? "" : i.ToString())
                .Subscribe(s => settings.UI.LookoutCoolDown.text = s)
                .AddTo(disposables);

            playerSubmarine.Resources.Torpedos.ForEach((torpedo, i) =>
            {
                var image = settings.UI.TorpedoResourceImages[i];
                torpedo.CanUse
                    .Subscribe(b =>
                    {
                        var color = Color.white;
                        color.a = b ? 1f : 0.35f;
                        image.color = color;
                    })
                    .AddTo(disposables);
            });
        }

        public void Dispose()
        {
            battleService.ResultDecided -= OnResultDecided;
            disposables.Dispose();
            battleService.FinishBattle();
        }

        public void Tick()
        {
            MoveToTitleUnlessInRoom();
            UpdateAlert();
            UpdateTimerText();
            UpdateDebugText();
        }

        void UpdateAlert()
        {
            var isActivePingerAlert = objectContainer.Submarines
                .OfType<EnemySubmarine>()
                .Any(s => s.IsUsingPinger);
            settings.UI.PingerAlert.gameObject.SetActive(isActivePingerAlert);
        }

        void UpdateTimerText()
        {
            var elapsedTimeSpan = DateTime.Now - battleService.StartDateTime;
            settings.UI.TimerText.text = string.Format(
                "{0:00}:{1:00}",
                elapsedTimeSpan.TotalMinutes,
                elapsedTimeSpan.Seconds
            );
        }

        void UpdateDebugText()
        {
            settings.UI.BattleLogText.text = string.Format(
                "{0:0.0} FPS\nPlayerId: {1:D} ({2})",
                Constants.Fps,
                connection.Player.ID,
                connection.Player.isMasterClient ? "Master" : "Slave"
            );
        }

        void MoveToTitleUnlessInRoom()
        {
            if (!connection.InRoom)
            {
                Debug.Log("Not in room");
                MoveToTitle();
            }
        }

        void MoveToTitle()
        {
            ZenUtil.LoadScene(Constants.SceneNames.Title);
        }

        void OnResultDecided(bool result)
        {
            if (result)
            {
                settings.UI.Victory.gameObject.SetActive(true);
            }
            else
            {
                settings.UI.Defeat.gameObject.SetActive(true);
            }
        }
    }
}
