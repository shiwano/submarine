using UnityEngine;
using System;
using UniRx;
using Zenject;
using Zenject.Commands;

namespace Submarine.Title
{
    public class SignUpCommand : Command<string>
    {
        public class Handler : ICommandHandler<string>
        {
            [Inject]
            UserModel user;
            [Inject]
            LobbyModel lobbyModel;
            [Inject]
            AuthenticationService auth;
            [Inject]
            PermanentDataStoreService dataStore;

            public void Execute(string userName)
            {
                if (dataStore.HasLoginData)
                {
                    throw new InvalidOperationException("Already signed to the API server.");
                }

                var password = Guid.NewGuid().ToString();

                auth.SignUp(userName, password).Subscribe(loggedInUser =>
                {
                    dataStore.UserName = userName;
                    dataStore.Password = password;
                    dataStore.Save();

                    lobbyModel.JoinedRoom.Value = loggedInUser.JoinedRoom;
                    user.LoggedInUser.Value = loggedInUser;
                    Debug.Log("Succeeded sign up");
                });
            }
        }
    }
}
