using UnityEngine;
using System;
using System.Collections;
using Zenject;
 
namespace Submarine
{
    public class MatchingService : MonoBehaviour
    {
        public event Action JoinedRoom;

        ConnectionService connection;

        [PostInject]
        public void Initialize(ConnectionService connection)
        {
            this.connection = connection;
        }

        public void JoinRoom()
        {
            connection.Connect();
        }

        void OnJoinedLobby()
        {
            Debug.Log("Joined Lobby");
            PhotonNetwork.JoinOrCreateRoom("Test", new RoomOptions(), TypedLobby.Default);
        }

        void OnJoinedRoom()
        {
            Debug.Log("Joined Room");
            connection.IsMessageQueueRunning = false;

            if (JoinedRoom != null)
            {
                JoinedRoom.Invoke();
            }
        }
    }
}