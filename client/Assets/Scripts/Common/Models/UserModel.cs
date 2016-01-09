using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class UserModel
    {
        public Type.LoggedInUser LoggedInUser { get; set; }

        public string Name { get { return LoggedInUser.Name; } }
        public bool IsInBattle { get { return LoggedInUser.JoinedRoom != null; } }
    }
}
