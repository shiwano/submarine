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
            [Inject]
            TyphenApi.WebApi.Submarine webApi;

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

                    webApi.Authenticate(response.AccessToken);
                    Debug.Log("Succeeded sign up");
                });
            }
        }
    }
}
