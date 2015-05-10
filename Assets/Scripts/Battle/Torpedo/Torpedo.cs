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
        readonly BattleService battleService;
        readonly BattleObjectContainer objectContainer;

        const float sqrSearchRange = 35f * 35f;

        public float LifeTime { get { return 6f; } }
        public Vector3 Acceleration { get { return Hooks.transform.forward * 20f; } }
        public Vector3 ShockPower { get { return Hooks.transform.forward * 40f; } }

        public PlayerTorpedo(TorpedoHooks hooks, BattleService battleService,
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
            var nearestEnemy = FindNearestTarget();
            Hooks.Target = nearestEnemy != null ? nearestEnemy.BattleObjectHooks.transform : null;
            Hooks.Accelerate(Acceleration * Constants.FpsRate);
        }

        IBattleObject FindNearestTarget()
        {
            var pair = objectContainer.Submarines
                .OfType<EnemySubmarine>()
                .Select(s => new { Source = s, SqrMagnitude = (Position - s.Position).sqrMagnitude })
                .Where(s => s.SqrMagnitude <= sqrSearchRange)
                .OrderBy(s => s.SqrMagnitude)
                .FirstOrDefault();
            return pair == null ? null : pair.Source;
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
