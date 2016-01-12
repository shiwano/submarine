using UnityEngine;
using System;
using System.Linq;
using UniRx;

namespace Submarine
{
    public interface ITorpedo : IBattleObject
    {
        TorpedoHooks Hooks { get; }
    }

    public abstract class TorpedoBase : ITorpedo
    {
        public TorpedoHooks Hooks { get; private set; }

        public BattleObjectType Type { get { return BattleObjectType.Torpedo; } }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }
        public Vector3 Position { get { return Hooks.transform.position; } }
        public Vector3 EulerAngles { get { return Hooks.transform.rotation.eulerAngles; } }
        public float SearchRange { get { return 35f; } }
        public bool IsVisibleFromPlayer { get { return true; } }

        protected TorpedoBase(TorpedoHooks hooks)
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

    public class PlayerTorpedo : TorpedoBase
    {
        readonly OldBattleService battleService;
        readonly BattleObjectContainer objectContainer;

        public float LifeTime { get { return 6f; } }
        public Vector3 Acceleration { get { return Hooks.transform.forward * 17f; } }
        public Vector3 ShockPower { get { return Hooks.transform.forward * 40f; } }

        public PlayerTorpedo(TorpedoHooks hooks, OldBattleService battleService,
            BattleObjectContainer objectContainer) : base(hooks)
        {
            this.battleService = battleService;
            this.objectContainer = objectContainer;
            Hooks.Striked += OnStriked;
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
            var nearestObject = FindNearestObjectInSearchRange();

            if (nearestObject != null && nearestObject.Type == BattleObjectType.Decoy)
            {
                Explode();
            }
            else
            {
                Hooks.Target = nearestObject != null ?
                    nearestObject.BattleObjectHooks.transform : null;
                Hooks.Accelerate(Acceleration * Constants.FpsRate);
            }
        }

        IBattleObject FindNearestObjectInSearchRange()
        {
            var decoy = objectContainer.FindNearestObjectInRange<EnemyDecoy>(Position, SearchRange);
            return decoy != null ?
                decoy as IBattleObject :
                objectContainer.FindNearestObjectInRange<EnemySubmarine>(Position, SearchRange);
        }

        void OnStriked(int? enemySubmarineViewId)
        {
            if (enemySubmarineViewId.HasValue)
            {
                battleService.SendSubmarineDamageEvent(
                    enemySubmarineViewId.Value,
                    Hooks.photonView.ownerId,
                    ShockPower
                );
            }
            Explode();
        }

        void Explode()
        {
            battleService.SendEffectPlayEvent(
                Constants.ExplosionEffectPrefab,
                Hooks.transform.position
            );
            objectContainer.Remove(this);
        }
    }

    public class EnemyTorpedo : TorpedoBase
    {
        public EnemyTorpedo(TorpedoHooks hooks) : base(hooks)
        {
        }
    }
}
