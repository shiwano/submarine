using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public abstract class ActorFacade : Facade
    {
        [Inject]
        Type.Battle.Actor actor;
        [Inject]
        UserModel userModel;

        public Type.Battle.Actor Actor { get { return actor; } }

        public bool IsMine
        {
            get { return actor.UserId == userModel.LoggedInUser.Value.Id; }
        }
    }
}
