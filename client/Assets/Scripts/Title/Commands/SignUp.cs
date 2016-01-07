using System;
using UniRx;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class SignUp : Command<string>
    {
        public class Handler : ICommandHandler<string>
        {
            readonly Models.User user;
            readonly Services.Authentication auth;
            readonly Services.PermanentDataStore dataStore;
            readonly Events.Title events;

            public Handler(
                Models.User user,
                Services.Authentication auth,
                Services.PermanentDataStore dataStore,
                Events.Title events)
            {
                this.user = user;
                this.auth = auth;
                this.dataStore = dataStore;
                this.events = events;
            }

            public void Execute(string userName)
            {
                if (dataStore.HasLoginData)
                {
                    throw new InvalidOperationException("Already signed to the API server.");
                }

                var password = Guid.NewGuid().ToString();

                auth.SignUp(userName, password).Subscribe(res =>
                {
                    dataStore.UserName = userName;
                    dataStore.Password = Guid.NewGuid().ToString();
                    dataStore.Save();

                    user.LoggedInUser = res.LoggedInUser;
                    user.ApiSessionKey = res.SessionKey;
                    events.LoginSucceeded.Invoke();
                });
            }
        }
    }
}
