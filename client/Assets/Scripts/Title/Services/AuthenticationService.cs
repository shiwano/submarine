using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Title
{
    public class AuthenticationService
    {
        [Inject]
        TyphenApi.WebApi.Submarine webApi;

        public IObservable<Type.LoggedInUser> Login(string userName, string passWord)
        {
            return webApi.Login(userName, passWord)
                .Send()
                .Select(r => r.Data.User);
        }

        public IObservable<Type.LoggedInUser> SignUp(string userName, string passWord)
        {
            return webApi.SignUp(userName, passWord)
                .Send()
                .Select(r => r.Data.User);
        }
    }
}
