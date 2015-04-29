using UnityEngine;
using System;
using System.Collections;
using Zenject;
 
namespace Submarine
{
    public class MatchingService : MonoBehaviour
    {
        public event Action onJoinRoom;

        private ConnectionService connection;

        [PostInject]
        public void Initialize(ConnectionService connection)
        {
            this.connection = connection;
        }

        public void JoinRoom()
        {
            connection.Connect();
        }

        private void OnJoinedLobby()
        {
            Debug.Log("Joined Lobby");
            PhotonNetwork.JoinOrCreateRoom("Test", new RoomOptions(), TypedLobby.Default);
        }

        private void OnJoinedRoom()
        {
            Debug.Log("Joined Room");
            connection.IsMessageQueueRunning = false;

            if (onJoinRoom != null)
            {
                onJoinRoom.Invoke();
            }
        }
    }
}