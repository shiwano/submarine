using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public abstract class ActorFacade : Facade
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
        public ActorMotor Motor { get { return motor; } }
        public ActorView View { get { return view; } }
        public bool IsMine { get { return actor.UserId == userModel.LoggedInUser.Value.Id; } }

        public override void Tick()
        {
            base.Tick();
            view.ActorPosition = motor.GetCurrentPosition();
        }
    }
}
