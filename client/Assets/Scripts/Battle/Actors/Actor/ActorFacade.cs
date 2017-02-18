using System;
using Zenject;
using Type = TyphenApi.Type.Submarine;
using UniRx;

namespace Submarine.Battle
{
    public abstract class ActorFacade : IInitializable, ITickable, IDisposable
    {
        [Inject]
        UserModel userModel;
        [Inject]
        Type.Battle.Actor actor;
        [Inject]
        ActorMotor motor;
        [Inject]
        ActorView view;

        public Type.Battle.Actor Actor { get { return actor; } }
        public ActorView View { get { return view; } }
        public bool IsMine { get { return actor.UserId == userModel.LoggedInUser.Value.Id; } }

        public virtual bool WillIgnoreMotorDirection { get { return false; } }

        public void Initialize()
        {
            motor.SetMovement(Actor.Movement);
            UpdatePositionAndDirection();

            if (!IsMine)
            {
                view.ChangeToEnemyColor();
            }

            Actor.OnMoveAsObservable().Subscribe(motor.SetMovement).AddTo(view);
        }

        public void Tick()
        {
            UpdatePositionAndDirection();
        }

        public void Dispose()
        {
            view.Dispose();
        }

        void UpdatePositionAndDirection()
        {
            view.ActorPosition = motor.GetCurrentPosition();

            if (!WillIgnoreMotorDirection)
            {
                view.ActorDirection = motor.GetCurrentDirection();
            }
        }
    }
}
