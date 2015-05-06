using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine
{
    public interface ITorpedo : IBattleObject
    {
        TorpedoHooks Hooks { get; }
    }

    public abstract class TorpedoBase : ITorpedo
    {
        public BattleObjectType Type { get { return BattleObjectType.Torpedo; } }
        public TorpedoHooks Hooks { get; private set; }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }

        protected TorpedoBase(TorpedoHooks hooks)
        {
            Hooks = hooks;
        }

        public virtual void Initialize() {}
        public virtual void Dispose() {}
        public virtual void Tick() {}
    }

    public class PlayerTorpedo : TorpedoBase
    {
        private readonly BattleService battleService;

        public Vector3 ShockPower { get { return Hooks.transform.forward * 300f; } }

        public Vector3 Speed
        {
            get { return Hooks.transform.forward * 50f; }
        }

        public PlayerTorpedo(TorpedoHooks hooks, BattleService battleService)
            : base(hooks)
        {
            this.battleService = battleService;
            Hooks.OnHitEnemySubmarine += OnHitEnemySubmarine;
        }

        public override void Tick()
        {
            Hooks.Accelerate(Speed);
        }

        private void OnHitEnemySubmarine(int enemySubmarineViewId)
        {
            battleService.SendSubmarineDamageEvent(
                enemySubmarineViewId,
                Hooks.photonView.ownerId,
                ShockPower
            );
        }
    }

    public class EnemyTorpedo : TorpedoBase
    {
        public EnemyTorpedo(TorpedoHooks hooks) : base(hooks)
        {
        }
    }
}
