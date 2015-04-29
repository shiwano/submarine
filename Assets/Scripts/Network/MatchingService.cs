using UnityEngine;
using System;
using System.Collections;
using Zenject;
 
namespace Submarine
{
    public class MatchingService : MonoBehaviour
    {
        public event Action OnJoinRoom;

        public bool IsMessageQueueRunning
        {
            get { return PhotonNetwork.isMessageQueueRunning; }
            set { PhotonNetwork.isMessageQueueRunning = value; }
        }

        public void Connect()
        {
            PhotonNetwork.ConnectUsingSettings(Constants.Version);
        }

        public void Disconnect()
        {
            PhotonNetwork.Disconnect();
        }

        /// <summary>
        /// PhotonNetwork.ConnectUsingSettings() でロビー接続が完了したときによばれる
        /// </summary>
        private void OnJoinedLobby()
        {
            Debug.Log("Joined Lobby");
            PhotonNetwork.JoinRandomRoom();
        }

        /// <summary>
        /// PhotonNetwork.JoinRandomRoom() でルームがなかったときによばれる
        /// </summary>
        private void OnPhotonRandomJoinFailed()
        {
            Debug.Log("Randam Join Failed");

            // TODO: ルーム作成機能を入れるまでは、ルーム名は仮
            PhotonNetwork.CreateRoom("Test");
        }

        /// <summary>
        /// ルームに入室したときによばれる
        /// </summary>
        private void OnJoinedRoom()
        {
            Debug.Log("Joined Room");
            IsMessageQueueRunning = false;

            if (OnJoinRoom != null)
            {
                OnJoinRoom.Invoke();
            }
        }
    }
}