using UnityEngine;
using System;
using UniRx;
using Zenject;
using Zenject.Commands;

namespace Submarine.Title
{
    public class LoginCommand : Command
    {
        public class Handler : ICommandHandler
        {
            [Inject]
            UserModel user;
            [Inject]
            AuthenticationService auth;
            [Inject]
            PermanentDataStoreService dataStore;
            [Inject]
            TitleEvents events;

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
