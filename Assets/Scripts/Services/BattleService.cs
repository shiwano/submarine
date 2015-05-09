using UnityEngine;
using System.Linq;
using Zenject;

namespace Submarine
{
    public class BattleService : Photon.MonoBehaviour
    {
        ConnectionService connection;
        BattleObjectContainer objectContainer;

        [PostInject]
        public void Initialize(ConnectionService connection, BattleObjectContainer objectContainer)
        {
            this.connection = connection;
            this.objectContainer = objectContainer;
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

        public void SendSubmarineDamageEvent(int damagedViewId, int attackerOwnerId, Vector3 shockPower)
        {
            photonView.RPC("ReceiveSubmarineDamageEvent", PhotonTargets.All, damagedViewId, attackerOwnerId, shockPower);
        }

        [RPC]
        void ReceiveSubmarineDamageEvent(int damagedViewId, int attackerOwnerId, Vector3 shockPower)
        {
            var damaged = objectContainer.Submarines.FirstOrDefault(s => s.Hooks.photonView.viewID == damagedViewId);
            var attacker = objectContainer.Submarines.FirstOrDefault(s => s.Hooks.photonView.ownerId == attackerOwnerId);
            BattleEvent.SubmarineDamaged(damaged, attacker, shockPower);
        }

        public void SendEffectPlayEvent(string resourceName, Vector3 position)
        {
            photonView.RPC("ReceiveEffectPlayEvent", PhotonTargets.All, resourceName, position);
        }

        [RPC]
        public void ReceiveEffectPlayEvent(string resourcePath, Vector3 position)
        {
            var prefab = Resources.Load(resourcePath);
            var effect = Instantiate(prefab) as GameObject;
            effect.transform.position = position;
        }
    }
}
