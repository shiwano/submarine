using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable, ITickable
    {
        readonly BattleInstaller.Settings settings;
        readonly BattleService battleService;
        readonly BattleObjectContainer objectContainer;
        readonly ThirdPersonCamera thirdPersonCamera;

        public BattleController(
            BattleInstaller.Settings settings,
            BattleService battleService,
            BattleObjectContainer objectContainer,
            ThirdPersonCamera thirdPersonCamera)
        {
            this.settings = settings;
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
            battleService.FinishBattle();
        }

        public void Tick()
        {
            UpdateTimerText();
            UpdateDebugText();
        }

        void UpdateTimerText()
        {
            var elapsedTimeSpan = DateTime.Now - battleService.StartDateTime;
            settings.UI.TimerText.text =
                elapsedTimeSpan.TotalMinutes.ToString("D2") + ":" +
                elapsedTimeSpan.Seconds.ToString("D2");
        }

        void UpdateDebugText()
        {
            settings.UI.BattleLogText.text = Constants.Fps.ToString("0.0") + " FPS";
        }

        void OnSubmarineDamaged(ISubmarine damaged, ISubmarine attacker, Vector3 shockPower)
        {
            damaged.Damage(shockPower);
        }
    }
}
