using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class UserModel
    {
        public readonly ReactiveProperty<Type.LoggedInUser> LoggedInUser;

        public readonly ReactiveProperty<string> Name;
        public readonly ReactiveProperty<bool> IsInBattle;

        public UserModel()
        {
            LoggedInUser = new ReactiveProperty<Type.LoggedInUser>();

            Name = LoggedInUser.Select(u => u == null ? string.Empty : u.Name).ToReactiveProperty();
            IsInBattle = LoggedInUser.Select(u => u != null && u.JoinedRoom != null).ToReactiveProperty();
        }
    }
}
