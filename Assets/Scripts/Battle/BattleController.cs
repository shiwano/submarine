using UnityEngine;
using System;
using System.Collections.Generic;
using System.Linq;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable
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

            var playerSubmarine = objectContainer.SpawnSubmarine(settings.Submarine.StartPositions[0]);
            thirdPersonCamera.SetTarget(playerSubmarine.Hooks.transform);
        }

        public void Dispose()
        {
            battleService.FinishBattle();
        }

        void OnSubmarineDamaged(ISubmarine sinked, ISubmarine attacker, Vector3 shockPower)
        {
            sinked.Damage(shockPower);
        }
    }
}
