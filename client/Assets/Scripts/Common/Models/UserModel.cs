using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class UserModel
    {
        public readonly ReactiveProperty<Type.LoggedInUser> LoggedInUser;

        public readonly ReadOnlyReactiveProperty<string> Name;

        public UserModel()
        {
            LoggedInUser = new ReactiveProperty<Type.LoggedInUser>();

            Name = LoggedInUser.Select(u => u == null ? string.Empty : u.Name).ToReadOnlyReactiveProperty();
        }
    }
}
