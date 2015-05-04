using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class BattleService : MonoBehaviour
    {
        private ConnectionService connection;

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
        }

        public void FinishBattle()
        {
            if (connection.InRoom)
            {
                PhotonNetwork.LeaveRoom();
            }
        }

        public GameObject InstantiatePhotonView(string prefabName, Vector3 position, Quaternion rotation)
        {
            return PhotonNetwork.Instantiate(prefabName, position, rotation, 0);
        }
    }
}
