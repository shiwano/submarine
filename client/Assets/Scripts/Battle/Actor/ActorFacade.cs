using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public interface IActorFacade : IFacade
    {
        Type.Battle.Actor Actor { get; }
        bool IsMine { get; }
    }

    public abstract class ActorFacade : Facade, IActorFacade
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
