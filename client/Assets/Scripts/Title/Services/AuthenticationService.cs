using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Title
{
    public class AuthenticationService
    {
        [Inject]
        TyphenApi.WebApi.Submarine webApi;

        public IObservable<Type.LoginObject> Login(string authToken)
        {
            return webApi.Login(authToken).Send().Select(x => x.Data);
        }

        public IObservable<Type.SignUpObject> SignUp(string name)
        {
            return webApi.SignUp(name).Send().Select(x => x.Data);
        }
    }
}
