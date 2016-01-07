using System;
using UniRx;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class Login : Command
    {
        public class Handler : ICommandHandler
        {
            readonly Models.User user;
            readonly Services.Authentication auth;
            readonly Services.PermanentDataStore dataStore;
            readonly TitleEvent titleEvent;

            public Handler(
                Models.User user,
                Services.Authentication auth,
                Services.PermanentDataStore dataStore,
                TitleEvent titleEvent)
            {
                this.user = user;
                this.auth = auth;
                this.dataStore = dataStore;
                this.titleEvent = titleEvent;
            }

            public void Execute()
            {
                if (!dataStore.HasLoginData)
                {
                    throw new InvalidOperationException("Not signed to the API server.");
                }

                auth.Login(dataStore.UserName, dataStore.Password).Subscribe(res =>
                {
                    user.LoggedInUser = res.LoggedInUser;
                    user.ApiSessionKey = res.SessionKey;
                    titleEvent.LoginSucceeded.Invoke();
                });
            }
        }
    }
}
