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
            LobbyModel lobbyModel;
            [Inject]
            AuthenticationService auth;
            [Inject]
            PermanentDataStoreService dataStore;

            public void Execute()
            {
                if (!dataStore.HasLoginData)
                {
                    throw new InvalidOperationException("Not signed to the API server.");
                }

                auth.Login(dataStore.UserName, dataStore.Password).Subscribe(loggedInUser =>
                {
                    user.LoggedInUser.Value = loggedInUser;
                    lobbyModel.JoinedRoom.Value = loggedInUser.JoinedRoom;
                    Debug.Log("Succeeded login");
                });
            }
        }
    }
}
