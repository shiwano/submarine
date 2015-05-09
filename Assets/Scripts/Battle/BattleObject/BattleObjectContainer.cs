using UnityEngine;
using System;
using System.Collections.Generic;
using System.Linq;
using Zenject;

namespace Submarine
{
    public class BattleObjectContainer : IDisposable, ITickable
    {
        readonly SubmarineFactory submarineFactory;
        readonly TorpedoFactory torpedoFactory;

        readonly List<IBattleObject> battleObjects = new List<IBattleObject>();

        public IEnumerable<ISubmarine> Submarines { get { return battleObjects.OfType<ISubmarine>(); } }
        public IEnumerable<ITorpedo> Torpedos { get { return battleObjects.OfType<ITorpedo>(); } }

        public event Action<IBattleObject> BattleObjectSpawned = delegate {};
        public event Action<IBattleObject> BattleObjectRemoved = delegate {};

        public BattleObjectContainer(
            SubmarineFactory submarineFactory,
            TorpedoFactory torpedoFactory)
        {
            this.submarineFactory = submarineFactory;
            this.torpedoFactory = torpedoFactory;

            BattleObjectHooksBase.CreatedViaNetwork += OnBattleObjectHooksCreatedViaNetwork;
            BattleObjectHooksBase.DestroyedViaNetwork += OnBattleObjectHooksDestroyedViaNetwork;
        }

        public void Dispose()
        {
            BattleObjectHooksBase.CreatedViaNetwork -= OnBattleObjectHooksCreatedViaNetwork;
            BattleObjectHooksBase.DestroyedViaNetwork -= OnBattleObjectHooksDestroyedViaNetwork;
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
            Add(submarine);
            return submarine;
        }

        public ITorpedo SpawnTorpedo(Vector3 position, Quaternion rotation)
        {
            var torpedo = torpedoFactory.Create(position, rotation);
            Add(torpedo);
            return torpedo;
        }

        public void Remove(IBattleObject battleObject)
        {
            var result = battleObjects.Remove(battleObject);
            if (result)
            {
                BattleObjectRemoved(battleObject);
                battleObject.Dispose();
            }
        }

        void Add(IBattleObject battleObject)
        {
            battleObjects.Add(battleObject);
            battleObject.Initialize();
            BattleObjectSpawned(battleObject);
        }

        void OnBattleObjectHooksCreatedViaNetwork(IBattleObjectHooks battleObjectHooks)
        {
            switch (battleObjectHooks.Type)
            {
                case BattleObjectType.Submarine:
                    var submarine = submarineFactory.Create(battleObjectHooks as SubmarineHooks);
                    Add(submarine);
                    break;
                case BattleObjectType.Torpedo:
                    var torpedo = torpedoFactory.Create(battleObjectHooks as TorpedoHooks);
                    Add(torpedo);
                    break;
            }
        }

        void OnBattleObjectHooksDestroyedViaNetwork(IBattleObjectHooks battleObjectHooks)
        {
            var battleObject = battleObjects.Find(s => s.BattleObjectHooks == battleObjectHooks);
            Remove(battleObject);
        }
    }
}
