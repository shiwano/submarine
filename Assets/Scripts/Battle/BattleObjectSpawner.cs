using UnityEngine;
using System;
using System.Collections.Generic;
using System.Linq;
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

        private void OnPhotonBehaviourCreate(IBattleObjectHooks battleObjectHooks)
        {
            if (battleObjects.Any(b => b.BattleObjectHooks == battleObjectHooks))
            {
                return;
            }

            switch (battleObjectHooks.Type)
            {
                case BattleObjectType.Submarine:
                    var submarine = submarineFactory.Create(battleObjectHooks as SubmarineHooks);
                    submarine.Initialize();
                    battleObjects.Add(submarine);
                    break;
                case BattleObjectType.Torpedo:
                    var torpedo = torpedoFactory.Create(battleObjectHooks as TorpedoHooks);
                    torpedo.Initialize();
                    battleObjects.Add(torpedo);
                    break;
            }
        }

        private void OnPhotonBehaviourDestroy(IBattleObjectHooks battleObjectHooks)
        {
            var battleObject = battleObjects.Find(s => s.BattleObjectHooks == battleObjectHooks);
            Destroy(battleObject);
        }
    }
}
