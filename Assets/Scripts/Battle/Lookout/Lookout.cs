using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    public interface ILookout : IBattleObject
    {
        LookoutHooks Hooks { get; }
    }

    public abstract class LookoutBase : ILookout
    {
        public LookoutHooks Hooks { get; private set; }

        public BattleObjectType Type { get { return BattleObjectType.Lookout; } }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }
        public Vector3 Position { get { return Hooks.transform.position; } }
        public Vector3 EulerAngles { get { return Hooks.transform.rotation.eulerAngles; } }
        public bool IsVisibleFromPlayer { get { return true; } }

        protected LookoutBase(LookoutHooks hooks)
        {
            Hooks = hooks;
        }

        public virtual void Initialize() {}
        public virtual void Tick() {}

        public virtual void Dispose()
        {
            Hooks.Dispose();
        }
    }

    public class PlayerLookout : LookoutBase
    {
        readonly BattleObjectContainer objectContainer;

        public float LifeTime { get { return 60f; } }

        public PlayerLookout(LookoutHooks hooks, BattleObjectContainer objectContainer) : base(hooks)
        {
            this.objectContainer = objectContainer;
        }

        public override void Initialize()
        {
            Observable.Interval(TimeSpan.FromSeconds(LifeTime))
                .Take(1)
                .Where(_ => Hooks != null)
                .Subscribe(_ => objectContainer.Remove(this));
        }
    }

    public class EnemyLookout : LookoutBase
    {
        public EnemyLookout(LookoutHooks hooks) : base(hooks)
        {
        }
    }
}
