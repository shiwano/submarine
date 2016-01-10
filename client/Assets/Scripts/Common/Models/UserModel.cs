using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class UserModel
    {
        public readonly ReactiveProperty<Type.LoggedInUser> LoggedInUser;

        public readonly ReadOnlyReactiveProperty<string> Name;
        public readonly ReadOnlyReactiveProperty<bool> IsInBattle;

        public UserModel()
        {
            LoggedInUser = new ReactiveProperty<Type.LoggedInUser>();

            Name = LoggedInUser.Select(u => u == null ? string.Empty : u.Name).ToReadOnlyReactiveProperty();
            IsInBattle = LoggedInUser.Select(u => u != null && u.JoinedRoom != null).ToReadOnlyReactiveProperty();
        }
    }
}
