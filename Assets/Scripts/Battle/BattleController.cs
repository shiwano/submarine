using UnityEngine;
using System;
using System.Collections.Generic;
using System.Linq;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable
    {
        private readonly BattleInstaller.Settings settings;
        private readonly BattleService battleService;
        private readonly BattleObjectContainer objectContainer;
        private readonly ThirdPersonCamera thirdPersonCamera;

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
            BattleEvent.OnSubmarineDamage += OnSubmarinDamage;

            battleService.StartBattle();

            var playerSubmarine = objectContainer.SpawnSubmarine(settings.Submarine.StartPositions[0]);
            thirdPersonCamera.SetTarget(playerSubmarine.Hooks.transform);
        }

        public void Dispose()
        {
            battleService.FinishBattle();
        }

        private void OnSubmarinDamage(ISubmarine sinked, ISubmarine attacker, Vector3 shockPower)
        {
            sinked.Damage(shockPower);
        }
    }
}
