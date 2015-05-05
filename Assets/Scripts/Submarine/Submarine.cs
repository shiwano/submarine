using UnityEngine;
using UniRx;
using Zenject;

namespace Submarine
{
    public interface ISubmarine : IBattleObject
    {
        SubmarineHooks Hooks { get; }
    }

    public abstract class SubmarineBase : ISubmarine
    {
        public BattleObjectType Type { get { return BattleObjectType.Submarine; } }
        public SubmarineHooks Hooks { get; private set; }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }

        protected SubmarineBase(SubmarineHooks hooks)
        {
            Hooks = hooks;
        }

        public virtual void Initialize() {}
        public virtual void Dispose() {}
        public virtual void Tick() {}
    }

    public class PlayerSubmarine : SubmarineBase
    {
        private readonly BattleInput input;
        private readonly BattleObjectSpawner spawner;
        private readonly CompositeDisposable eventResources = new CompositeDisposable();

        public Vector3 Speed
        {
            get { return Hooks.transform.forward * Mathf.Lerp(0f, 50f, Mathf.Clamp01(input.MousePressingTime)); }
        }

        public Vector3 TurningEulerAngles
        {
            get { return Hooks.transform.up * (input.MousePosition.x - input.MousePositionOnButtonDown.x) * 0.01f; }
        }

        public PlayerSubmarine(SubmarineHooks hooks, BattleInput input, BattleObjectSpawner spawner)
            : base(hooks)
        {
            this.input = input;
            this.spawner = spawner;
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
            if (input.IsMouseButtonPressed.Value)
            {
                Hooks.Accelerate(Speed);
                Hooks.Turn(TurningEulerAngles);
            }
        }

        private void SpawnTorpedo()
        {
            spawner.SpawnTorpedo(Hooks.LaunchSitePosition, Hooks.transform.rotation);
        }
    }

    public class EnemySubmarine : SubmarineBase
    {
        public EnemySubmarine(SubmarineHooks hooks) : base(hooks)
        {
        }
    }
}
