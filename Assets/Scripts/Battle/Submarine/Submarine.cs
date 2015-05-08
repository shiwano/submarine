using UnityEngine;
using UniRx;
using Zenject;

namespace Submarine
{
    public interface ISubmarine : IBattleObject
    {
        SubmarineHooks Hooks { get; }
        void Damage(Vector3 shockPower);
    }

    public abstract class SubmarineBase : ISubmarine
    {
        public SubmarineHooks Hooks { get; private set; }

        public BattleObjectType Type { get { return BattleObjectType.Submarine; } }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }
        public Vector3 Position { get { return Hooks.transform.position; } }
        public Vector3 EulerAngles { get { return Hooks.transform.rotation.eulerAngles; } }

        protected SubmarineBase(SubmarineHooks hooks)
        {
            Hooks = hooks;
        }

        public virtual void Initialize() {}
        public virtual void Dispose() {}
        public virtual void Tick() {}

        public virtual void Damage(Vector3 shockPower)
        {
            Hooks.Damage(shockPower);
        }
    }

    public class PlayerSubmarine : SubmarineBase
    {
        readonly BattleInput input;
        readonly BattleObjectContainer objectContainer;
        readonly CompositeDisposable eventResources = new CompositeDisposable();

        bool IsSinked = false;

        public Vector3 Acceleration
        {
            get { return Hooks.transform.forward * Mathf.Lerp(0f, 60f, Mathf.Clamp01(input.MousePressingTime)); }
        }

        public Vector3 TurningEulerAngles
        {
            get { return Hooks.transform.up * (input.MousePosition.x - input.MousePositionOnButtonDown.x) * 0.01f; }
        }

        public PlayerSubmarine(SubmarineHooks hooks, BattleInput input, BattleObjectContainer objectContainer)
            : base(hooks)
        {
            this.input = input;
            this.objectContainer = objectContainer;
        }

        public override void Initialize()
        {
            input.IsMouseButtonPressed
                .Where(b => !b)
                .Subscribe(_ => Hooks.Brake())
                .AddTo(eventResources);

            input.IsMouseButtonClicked
                .Skip(1)
                .Where(b => b)
                .Subscribe(_ => SpawnTorpedo())
                .AddTo(eventResources);
        }

        public override void Dispose()
        {
            eventResources.Dispose();
        }
       
        public override void Tick()
        {
            if (!IsSinked && input.IsMouseButtonPressed.Value)
            {
                Hooks.Accelerate(Acceleration * Constants.FpsRate);
                Hooks.Turn(TurningEulerAngles);
            }
        }

        public override void Damage(Vector3 shockPower)
        {
            IsSinked = true;
            eventResources.Dispose();
            base.Damage(shockPower);
        }

        void SpawnTorpedo()
        {
            objectContainer.SpawnTorpedo(Hooks.LaunchSitePosition, Hooks.transform.rotation);
        }
    }

    public class EnemySubmarine : SubmarineBase
    {
        public EnemySubmarine(SubmarineHooks hooks) : base(hooks)
        {
        }
    }
}
