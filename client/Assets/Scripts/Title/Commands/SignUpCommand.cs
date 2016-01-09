using UnityEngine;
using System;
using UniRx;
using Zenject.Commands;

namespace Submarine.Title
{
    public class SignUpCommand : Command<string>
    {
        public class Handler : ICommandHandler<string>
        {
            readonly UserModel user;
            readonly AuthenticationService auth;
            readonly PermanentDataStoreService dataStore;
            readonly TitleEvents events;

            public Handler(
                UserModel user,
                AuthenticationService auth,
                PermanentDataStoreService dataStore,
                TitleEvents events)
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
                    dataStore.Password = password;
                    dataStore.Save();

                    user.LoggedInUser = res.LoggedInUser;
                    user.ApiSessionKey = res.SessionKey;
                    events.LoginSucceeded.Invoke();
                    Debug.Log("Succeeded sign up");
                });
            }
        }
    }
}
