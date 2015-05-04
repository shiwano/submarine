using UnityEngine;
using System;
using System.Collections.Generic;
using Zenject;

namespace Submarine
{
    public class BattleObjectSpawner : IInitializable, IDisposable, ITickable
    {
        private readonly SubmarineFactory submarineFactory;
        private readonly TorpedoFactory torpedoFactory;

        private readonly List<ISubmarine> submarines = new List<ISubmarine>();
        private readonly List<ITorpedo> torpedos = new List<ITorpedo>();

        public BattleObjectSpawner(
            SubmarineFactory submarineFactory,
            TorpedoFactory torpedoFactory)
        {
            this.submarineFactory = submarineFactory;
            this.torpedoFactory = torpedoFactory;
        }

        public void Initialize()
        {
            BattleEvent.OnPhotonBehaviourCreate += OnPhotonBehaviourCreate;
            BattleEvent.OnPhotonBehaviourDestroy += OnPhotonBehaviourDestroy;
        }

        public void Dispose()
        {
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

        public ISubmarine SpawnSubmarine(Vector3 position)
        {
            var submarine = submarineFactory.Create(position);
            submarine.Initialize();
            submarines.Add(submarine);
            return submarine;
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
                submarine.Initialize();
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
