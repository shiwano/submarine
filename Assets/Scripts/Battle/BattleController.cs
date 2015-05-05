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
        private readonly BattleObjectSpawner spawner;
        private readonly ThirdPersonCamera thirdPersonCamera;

        public BattleController(
            BattleInstaller.Settings settings,
            BattleService battleService,
            BattleObjectSpawner spawner,
            ThirdPersonCamera thirdPersonCamera)
        {
            this.settings = settings;
            this.battleService = battleService;
            this.spawner = spawner;
            this.thirdPersonCamera = thirdPersonCamera;
        }

        public void Initialize()
        {
            BattleEvent.OnSubmarineSink += OnSubmarinSink;

            battleService.StartBattle();

            var playerSubmarine = spawner.SpawnSubmarine(settings.Submarine.StartPositions[0]);
            thirdPersonCamera.SetTarget(playerSubmarine.Hooks.transform);
        }

        public void Dispose()
        {
            battleService.FinishBattle();
        }

        private void OnSubmarinSink(ISubmarine sinked, ISubmarine attacker)
        {
            sinked.Sink();
        }
    }
}
