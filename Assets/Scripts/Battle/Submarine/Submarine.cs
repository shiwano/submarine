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
        public static readonly float sqrSearchRange = 70f * 70f;

        public SubmarineHooks Hooks { get; private set; }
        public bool IsSinked { get; protected set; }
        public abstract bool IsUsingPinger { get; }

        public BattleObjectType Type { get { return BattleObjectType.Submarine; } }
        public IBattleObjectHooks BattleObjectHooks { get { return Hooks; } }
        public Vector3 Position { get { return Hooks.transform.position; } }
        public Vector3 EulerAngles { get { return Hooks.transform.rotation.eulerAngles; } }
        public virtual bool IsVisibleFromPlayer { get { return true; } }

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
            return sqrLength <= SubmarineBase.sqrSearchRange;
        }
    }

    public class PlayerSubmarine : SubmarineBase
    {
        readonly BattleInput input;
        readonly BattleObjectContainer objectContainer;
        readonly SubmarineResources resources;
        readonly BattleService battleService;

        readonly CompositeDisposable disposables = new CompositeDisposable();

        public override bool IsUsingPinger { get { return resources.Pinger.IsUsing.Value; } }
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
            SubmarineResources resources,
            BattleService battleService) : base(hooks)
        {
            this.input = input;
            this.objectContainer = objectContainer;
            this.resources = resources;
            this.battleService = battleService;
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
                .Subscribe(_ => Debug.Log("Decoy"))
                .AddTo(disposables);

            input.PingerButtonClickedAsObservable
                .Subscribe(_ => UsePinger())
                .AddTo(disposables);

            input.LookoutButtonClickedAsObservable
                .Subscribe(_ => Debug.Log("Lookout"))
                .AddTo(disposables);

            resources.Pinger.IsUsing
                .Skip(1)
                .Subscribe(b => battleService.SendPingerEvent(Hooks.ViewId, b))
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
            base.Damage(shockPower);
        }

        void UseTorpedo()
        {
            var usableTorpedo = resources.Torpedos.FirstOrDefault(t => t.CanUse.Value);
            if (usableTorpedo != null)
            {
                usableTorpedo.Use();
                objectContainer.SpawnTorpedo(Hooks.LaunchSitePosition, Hooks.transform.rotation);
            }
        }

        void UsePinger()
        {
            if (resources.Pinger.CanUse.Value)
            {
                resources.Pinger.Use();
                battleService.SendPingerEvent(Hooks.ViewId, IsUsingPinger);
            }
        }
    }

    public class EnemySubmarine : SubmarineBase
    {
        readonly BattleObjectContainer objectContainer;

        bool isUsingPinger;
        public override bool IsUsingPinger { get { return isUsingPinger; } }

        public EnemySubmarine(SubmarineHooks hooks, BattleObjectContainer objectContainer)
            : base(hooks)
        {
            this.objectContainer = objectContainer;
        }

        public void SetUsingPinger(bool isUsingPinger)
        {
            this.isUsingPinger = isUsingPinger;
        }

        public override bool IsVisibleFromPlayer
        {
            get
            {
                foreach (var playerSubmarine in objectContainer.Submarines.OfType<PlayerSubmarine>())
                {
                    if (playerSubmarine.Resources.Pinger.IsUsing.Value ||
                        IsInSearchRangeOf(playerSubmarine))
                    {
                        return true;
                    }
                }
                return false;
            }
        }
    }
}
