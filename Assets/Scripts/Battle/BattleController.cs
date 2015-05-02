using UnityEngine;
using System;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable
    {
        private readonly BattleInstaller.Settings settings;
        private readonly BattleService battleService;
        private readonly SubmarineFactory submarineFactory;
        private readonly ThirdPersonCamera thirdPersonCamera;

        public BattleController(
            BattleInstaller.Settings settings,
            BattleService battleService,
            SubmarineFactory submarineFactory,
            ThirdPersonCamera thirdPersonCamera)
        {
            this.settings = settings;
            this.battleService = battleService;
            this.submarineFactory = submarineFactory;
            this.thirdPersonCamera = thirdPersonCamera;
        }

        public void Initialize()
        {
            battleService.StartBattle();

            var playerSubmarine = submarineFactory.Create(settings.Submarine.StartPositions[0]);
            thirdPersonCamera.SetTarget(playerSubmarine.Transform);
        }

        public void Dispose()
        {
            battleService.FinishBattle();
        }
    }
}
