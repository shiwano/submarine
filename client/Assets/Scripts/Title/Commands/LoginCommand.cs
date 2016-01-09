using UnityEngine;
using System;
using UniRx;
using Zenject.Commands;

namespace Submarine.Title
{
    public class LoginCommand : Command
    {
        public class Handler : ICommandHandler
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
                    events.LoginSucceeded.Invoke();
                    Debug.Log("Succeeded login");
                });
            }
        }
    }
}
