using UnityEngine;
using System;
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

        public BattleController(
            BattleInstaller.Settings settings,
            ConnectionService connection,
            BattleService battleService,
            BattleObjectContainer objectContainer,
            ThirdPersonCamera thirdPersonCamera)
        {
            this.settings = settings;
            this.connection = connection;
            this.battleService = battleService;
            this.objectContainer = objectContainer;
            this.thirdPersonCamera = thirdPersonCamera;
        }

        public void Initialize()
        {
            BattleEvent.SubmarineDamaged += OnSubmarineDamaged;
            battleService.StartBattle();

            var playerSubmarine = objectContainer.SpawnSubmarine(settings.Map.StartPositions[0]);
            thirdPersonCamera.SetTarget(playerSubmarine.Hooks.transform);
        }

        public void Dispose()
        {
            BattleEvent.SubmarineDamaged -= OnSubmarineDamaged;
            battleService.FinishBattle();
        }

        public void Tick()
        {
            MoveToTitleUnlessInRoom();
            UpdateTimerText();
            UpdateDebugText();
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
                ZenUtil.LoadScene(Constants.SceneNames.Title);
            }
        }

        void OnSubmarineDamaged(ISubmarine damaged, ISubmarine attacker, Vector3 shockPower)
        {
            damaged.Damage(shockPower);
        }
    }
}
