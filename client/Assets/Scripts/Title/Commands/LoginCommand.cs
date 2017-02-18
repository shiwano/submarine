using UnityEngine;
using System;
using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class LoginCommand : Signal<LoginCommand>
    {
        public class Handler
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
                if (!dataStore.HasSignedUp)
                {
                    throw new InvalidOperationException("Not signed to the API server.");
                }

                auth.Login(dataStore.AuthToken).Subscribe(response =>
                {
                    lobbyModel.JoinedRoom.Value = response.User.JoinedRoom;
                    user.LoggedInUser.Value = response.User;
                    Debug.Log("Succeeded login");
                });
            }
        }
    }
}
