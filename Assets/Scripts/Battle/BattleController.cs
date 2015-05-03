using UnityEngine;
using System;
using System.Collections.Generic;
using System.Linq;
using Zenject;

namespace Submarine
{
    public class BattleController : IInitializable, IDisposable, ITickable
    {
        private readonly BattleInstaller.Settings settings;
        private readonly BattleService battleService;
        private readonly SubmarineFactory submarineFactory;
        private readonly ThirdPersonCamera thirdPersonCamera;

        private readonly List<ISubmarine> submarines = new List<ISubmarine>();

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
            BattleEvent.OnPhotonBehaviourCreate += OnPhotonBehaviourCreate;
            BattleEvent.OnPhotonBehaviourDestroy += OnPhotonBehaviourDestroy;
            battleService.StartBattle();

            var playerSubmarine = submarineFactory.Create(settings.Submarine.StartPositions[0]);
            submarines.Add(playerSubmarine);

            thirdPersonCamera.SetTarget(playerSubmarine.Hooks.transform);
        }

        public void Dispose()
        {
            battleService.FinishBattle();
            BattleEvent.OnPhotonBehaviourCreate -= OnPhotonBehaviourCreate;
            BattleEvent.OnPhotonBehaviourDestroy -= OnPhotonBehaviourDestroy;
        }

        public void Tick()
        {
            foreach (var submarine in submarines)
            {
                submarine.Tick();
            }
        }

        private void OnPhotonBehaviourCreate(Photon.MonoBehaviour photonMonoBehaviour)
        {
            if (photonMonoBehaviour.photonView.isMine)
            {
                return;
            }

            var submarineHooks = photonMonoBehaviour as SubmarineHooks;
            if (submarineHooks != null)
            {
                var submarine = submarineFactory.Create(submarineHooks);
                submarines.Add(submarine);
            }
        }

        private void OnPhotonBehaviourDestroy(Photon.MonoBehaviour photonMonoBehaviour)
        {
            var submarineHooks = photonMonoBehaviour as SubmarineHooks;
            if (submarineHooks != null)
            {
                var destroyedIndex = submarines.FindIndex(s => s.Hooks == submarineHooks);
                submarines.RemoveAt(destroyedIndex);
            }
        }
    }
}
