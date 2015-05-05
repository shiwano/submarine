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

        private readonly List<IBattleObject> battleObjects = new List<IBattleObject>();

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
            foreach (var battleObject in battleObjects)
            {
                battleObject.Tick();
            }
        }

        public ISubmarine SpawnSubmarine(Vector3 position)
        {
            var submarine = submarineFactory.Create(position);
            submarine.Initialize();
            battleObjects.Add(submarine);
            return submarine;
        }

        public ITorpedo SpawnTorpedo(Vector3 position, Quaternion rotation)
        {
            var torpedo = torpedoFactory.Create(position, rotation);
            torpedo.Initialize();
            battleObjects.Add(torpedo);
            return torpedo;
        }

        public void Destroy(IBattleObject battleObject)
        {
            var result = battleObjects.Remove(battleObject);
            if (result)
            {
                battleObject.Dispose();
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
                submarine.Initialize();
                battleObjects.Add(submarine);
            }

            var torpedoHooks = photonMonoBehaviour as TorpedoHooks;
            if (torpedoHooks != null)
            {
                var torpedo = torpedoFactory.Create(torpedoHooks);
                battleObjects.Add(torpedo);
            }
        }

        private void OnPhotonBehaviourDestroy(Photon.MonoBehaviour photonMonoBehaviour)
        {
            var battleObject = battleObjects.Find(s => s.PhotonMonoBehaviour == photonMonoBehaviour);
            Destroy(battleObject);
        }
    }
}
