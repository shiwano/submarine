using UnityEngine;
using System;
using System.Linq;
using Zenject;

namespace Submarine
{
    public class BattleService : Photon.MonoBehaviour
    {
        private ConnectionService connection;
        private BattleObjectSpawner spawner;

        [PostInject]
        public void Initialize(ConnectionService connection, BattleObjectSpawner spawner)
        {
            this.connection = connection;
            this.spawner = spawner;
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

        public void SendSubmarineSinkEvent(int sinkedViewId, int attackerOwnerId)
        {
            photonView.RPC("ReceiveSubmarineSinkEvent", PhotonTargets.All, sinkedViewId, attackerOwnerId);
        }

        [RPC]
        private void ReceiveSubmarineSinkEvent(int sinkedViewId, int attackerOwnerId)
        {
            var sinked = spawner.Submarines.FirstOrDefault(s => s.Hooks.photonView.viewID == sinkedViewId);
            var attacker = spawner.Submarines.FirstOrDefault(s => s.Hooks.photonView.ownerId == attackerOwnerId);
            BattleEvent.OnSubmarineSink(sinked, attacker);
        }

        public GameObject InstantiatePhotonView(string prefabName, Vector3 position, Quaternion rotation)
        {
            return PhotonNetwork.Instantiate(prefabName, position, rotation, 0);
        }
    }
}
