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
            [Inject]
            TyphenApi.WebApi.Submarine webApi;

            public void Execute()
            {
                if (!dataStore.HasSignedUp)
                {
                    throw new InvalidOperationException("Not signed to the API server.");
                }

                auth.Login(dataStore.AuthToken).Subscribe(response =>
                {
                    lobbyModel.JoinedRoom.Value = response.User.JoinedRoom;
                    user.LoggedInUser.Value = response.User;

                    webApi.Authenticate(response.AccessToken);
                    Debug.Log("Succeeded login");
                });
            }
        }
    }
}
