using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    public interface IDecoy : IBattleObject
    {
        DecoyHooks Hooks { get; }
    }

    public abstract class DecoyBase : IDecoy
    {
        public DecoyHooks Hooks { get; private set; }

        public BattleObjectType Type { get { return BattleObjectType.Decoy; } }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }
        public Vector3 Position { get { return Hooks.transform.position; } }
        public Vector3 EulerAngles { get { return Hooks.transform.rotation.eulerAngles; } }
        public bool IsVisibleFromPlayer { get { return true; } }

        protected DecoyBase(DecoyHooks hooks)
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

    public class PlayerDecoy : DecoyBase
    {
        readonly BattleObjectContainer objectContainer;

        public float LifeTime { get { return 5f; } }
        public Vector3 Acceleration { get { return Hooks.transform.forward * 5f; } }

        public PlayerDecoy(DecoyHooks hooks, BattleObjectContainer objectContainer) : base(hooks)
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

        public override void Tick()
        {
            Hooks.Accelerate(Acceleration * Constants.FpsRate);
        }
    }

    public class EnemyDecoy : DecoyBase
    {
        public EnemyDecoy(DecoyHooks hooks) : base(hooks)
        {
        }
    }
}
