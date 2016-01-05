using UnityEngine;
using System;
using System.Linq;
using UniRx;

namespace Submarine
{
    public interface ISubmarine : IBattleObject
    {
        SubmarineHooks Hooks { get; }
        bool IsSinked { get; }
        bool IsUsingPinger { get; }
        void Damage(Vector3 shockPower);
    }

    public abstract class SubmarineBase : ISubmarine
    {
        public SubmarineHooks Hooks { get; private set; }
        public bool IsSinked { get; protected set; }

        public BattleObjectType Type { get { return BattleObjectType.Submarine; } }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }
        public Vector3 Position { get { return Hooks.transform.position; } }
        public Vector3 EulerAngles { get { return Hooks.transform.rotation.eulerAngles; } }
        public float SearchRange { get { return 70f; } }
        public virtual bool IsVisibleFromPlayer { get { return true; } }
        public bool IsUsingPinger { get { return Hooks.IsUsingPinger; } }

        protected SubmarineBase(SubmarineHooks hooks)
        {
            Hooks = hooks;
        }

        public virtual void Initialize() {}
        public virtual void Tick() {}

        public virtual void Dispose()
        {
            Hooks.Dispose();
        }

        public virtual void Damage(Vector3 shockPower)
        {
            IsSinked = true;
            Hooks.Damage(shockPower);
        }

        public bool IsInSearchRangeOf(IBattleObject battleObject)
        {
            var sqrLength = (battleObject.Position - Position).sqrMagnitude;
            return sqrLength <= (battleObject.SearchRange * battleObject.SearchRange);
        }
    }

    public class PlayerSubmarine : SubmarineBase
    {
        readonly BattleInput input;
        readonly BattleObjectContainer objectContainer;
        readonly SubmarineResources resources;

        readonly CompositeDisposable disposables = new CompositeDisposable();

        public SubmarineResources Resources { get { return resources; } }

        public Vector3 Acceleration
        {
            get { return Hooks.transform.forward * Mathf.Lerp(0f, 6f, Mathf.Clamp01(input.TouchTime)); }
        }

        public Vector3 TurningEulerAngles
        {
            get { return Hooks.transform.up * input.DragAmount.x * 0.01f; }
        }

        public PlayerSubmarine(
            SubmarineHooks hooks,
            BattleInput input,
            BattleObjectContainer objectContainer,
            SubmarineResources resources) : base(hooks)
        {
            this.input = input;
            this.objectContainer = objectContainer;
            this.resources = resources;
        }

        public override void Initialize()
        {
            input.IsTouched
                .Where(b => !b)
                .Subscribe(_ => Hooks.Brake())
                .AddTo(disposables);

            input.ClickedAsObservable
                .Subscribe(_ => UseTorpedo())
                .AddTo(disposables);

            input.DecoyButtonClickedAsObservable
                .Subscribe(_ => UseDecoy())
                .AddTo(disposables);

            input.PingerButtonClickedAsObservable
                .Subscribe(_ => UsePinger())
                .AddTo(disposables);

            input.LookoutButtonClickedAsObservable
                .Subscribe(_ => UseLookout())
                .AddTo(disposables);

            resources.Pinger.IsUsing
                .Subscribe(b => Hooks.IsUsingPinger = b)
                .AddTo(disposables);
        }

        public override void Dispose()
        {
            disposables.Dispose();
            base.Dispose();
        }
       
        public override void Tick()
        {
            if (!IsSinked && input.IsTouched.Value)
            {
                Hooks.Accelerate(Acceleration * Constants.FpsRate);
                Hooks.Turn(TurningEulerAngles);
            }
        }

        public override void Damage(Vector3 shockPower)
        {
            IsSinked = true;
            disposables.Dispose();
            base.Damage(shockPower);
        }

        void UseTorpedo()
        {
            var usableTorpedo = resources.Torpedos.FirstOrDefault(t => t.CanUse.Value);
            if (usableTorpedo != null)
            {
                usableTorpedo.Use();
                objectContainer.SpawnTorpedo(
                    Hooks.TorpedoLaunchSitePosition,
                    Hooks.transform.rotation
                );
            }
        }

        void UseDecoy()
        {
            if (resources.Decoy.CanUse.Value)
            {
                resources.Decoy.Use();
                objectContainer.SpawnDecoy(
                    Hooks.DecoyLaunchSitePosition,
                    Hooks.DecoyLaunchSiteRotation
                );
            }
        }

        void UsePinger()
        {
            if (resources.Pinger.CanUse.Value)
            {
                resources.Pinger.Use();
            }
        }

        void UseLookout()
        {
            if (resources.Lookout.CanUse.Value)
            {
                resources.Lookout.Use();
                objectContainer.SpawnLookout(
                    Hooks.LookoutLaunchSitePosition,
                    Hooks.transform.rotation
                );
            }
        }
    }

    public class EnemySubmarine : SubmarineBase
    {
        readonly BattleObjectContainer objectContainer;

        public EnemySubmarine(SubmarineHooks hooks, BattleObjectContainer objectContainer)
            : base(hooks)
        {
            this.objectContainer = objectContainer;
        }

        public override bool IsVisibleFromPlayer
        {
            get
            {
                return
                    objectContainer.Submarines
                        .OfType<PlayerSubmarine>()
                        .Any(s => s.IsUsingPinger || IsInSearchRangeOf(s)) ||
                    objectContainer.Lookouts
                        .OfType<PlayerLookout>()
                        .Any(IsInSearchRangeOf);
            }
        }
    }
}
