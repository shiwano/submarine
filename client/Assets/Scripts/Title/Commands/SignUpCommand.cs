using UnityEngine;
using System;
using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class SignUpCommand : Signal<string, SignUpCommand>
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

            public void Execute(string userName)
            {
                if (dataStore.HasSignedUp)
                {
                    throw new InvalidOperationException("Already signed to the API server.");
                }

                auth.SignUp(userName).Subscribe(response =>
                {
                    dataStore.AuthToken = response.AuthToken;
                    dataStore.Save();

                    lobbyModel.JoinedRoom.Value = response.User.JoinedRoom;
                    user.LoggedInUser.Value = response.User;
                    Debug.Log("Succeeded sign up");
                });
            }
        }
    }
}
