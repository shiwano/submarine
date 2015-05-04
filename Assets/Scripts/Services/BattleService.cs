using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class BattleService : MonoBehaviour
    {
        private ConnectionService connection;

        public int PlayerNumber
        {
            get { return Array.FindIndex(PhotonNetwork.playerList, p => p == PhotonNetwork.player); }
        }

        [PostInject]
        public void Initialize(ConnectionService connection)
        {
            this.connection = connection;
        }

        public void StartBattle()
        {
            if (!connection.InRoom)
            {
                Debug.LogError("Not in room");
                return;
            }

            connection.IsMessageQueueRunning = true;
            Debug.Log(PlayerNumber);
        }

        public void FinishBattle()
        {
            if (connection.InRoom)
            {
                PhotonNetwork.LeaveRoom();
            }
        }

        public GameObject InstantiatePhotonView(string prefabName, Vector3 position)
        {
            return PhotonNetwork.Instantiate(prefabName, position, Quaternion.identity, 0);
        }
    }
}
