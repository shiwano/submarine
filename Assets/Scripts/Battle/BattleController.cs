using UnityEngine;
using System;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable
    {
        private readonly BattleService battleService;
        private readonly SubmarineFactory submarineFactory;

        public BattleController(BattleService battleService, SubmarineFactory submarineFactory)
        {
            this.battleService = battleService;
            this.submarineFactory = submarineFactory;
        }

        public void Initialize()
        {
            battleService.StartBattle();

            submarineFactory.Create(Vector3.zero);
        }

        public void Dispose()
        {
            battleService.FinishBattle();
        }
    }
}
