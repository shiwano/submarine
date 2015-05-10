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
        readonly DecoyFactory decoyFactory;

        readonly List<IBattleObject> battleObjects = new List<IBattleObject>();

        public IEnumerable<ISubmarine> Submarines { get { return battleObjects.OfType<ISubmarine>(); } }
        public IEnumerable<ISubmarine> AliveSubmarines { get { return Submarines.Where(s => !s.IsSinked); } }
        public IEnumerable<ITorpedo> Torpedos { get { return battleObjects.OfType<ITorpedo>(); } }

        public event Action<IBattleObject> BattleObjectSpawned = delegate {};
        public event Action<IBattleObject> BattleObjectRemoved = delegate {};

        public BattleObjectContainer(
            SubmarineFactory submarineFactory,
            TorpedoFactory torpedoFactory,
            DecoyFactory decoyFactory)
        {
            this.submarineFactory = submarineFactory;
            this.torpedoFactory = torpedoFactory;
            this.decoyFactory = decoyFactory;

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

        public PlayerSubmarine SpawnPlayerSubmarine(Vector3 position)
        {
            var submarine = submarineFactory.CreatePlayerSubmarine(position);
            Add(submarine);
            return submarine;
        }

        public ITorpedo SpawnTorpedo(Vector3 position, Quaternion rotation)
        {
            var torpedo = torpedoFactory.Create(position, rotation);
            Add(torpedo);
            return torpedo;
        }

        public IDecoy SpawnDecoy(Vector3 position, Quaternion rotation)
        {
            var decoy = decoyFactory.Create(position, rotation);
            Add(decoy);
            return decoy;
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
            if (battleObjectHooks.IsMine) { return; }

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
                case BattleObjectType.Decoy:
                    var decoy = decoyFactory.Create(battleObjectHooks as DecoyHooks);
                    Add(decoy);
                    break;
            }
        }

        void OnBattleObjectHooksDestroyedViaNetwork(IBattleObjectHooks battleObjectHooks)
        {
            var battleObject = battleObjects.Find(s => s.BattleObjectHooks == battleObjectHooks);

            if (battleObject != null)
            {
                Remove(battleObject);
            }
        }
    }
}
