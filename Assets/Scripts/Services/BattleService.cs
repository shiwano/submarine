using UnityEngine;
using System;
using System.Linq;
using Zenject;
using UniRx;

namespace Submarine
{
    public class BattleService : Photon.MonoBehaviour
    {
        ConnectionService connection;
        BattleObjectContainer objectContainer;

        public DateTime StartDateTime { get; private set; }
        public ReactiveProperty<int> EnemyPingerCounter { get; private set; }

        [PostInject]
        public void Initialize(ConnectionService connection, BattleObjectContainer objectContainer)
        {
            this.connection = connection;
            this.objectContainer = objectContainer;

            EnemyPingerCounter = new ReactiveProperty<int>(0);
        }

        public void StartBattle()
        {
            if (connection.InRoom)
            {
                StartDateTime = DateTime.Now;
                connection.IsMessageQueueRunning = true;

                if (connection.Player.isMasterClient)
                {
                    SendSynchronizeStartTimeEvent();
                }
            }
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

        public void DestroyPhotonView(GameObject targetView)
        {
            PhotonNetwork.Destroy(targetView);
        }

        public void SendSubmarineDamageEvent(int damagedViewId, int attackerOwnerId, Vector3 shockPower)
        {
            photonView.RPC("ReceiveSubmarineDamageEvent", PhotonTargets.All, damagedViewId, attackerOwnerId, shockPower);
        }

        [RPC]
        void ReceiveSubmarineDamageEvent(int damagedViewId, int attackerOwnerId, Vector3 shockPower)
        {
            var damaged = objectContainer.Submarines.FirstOrDefault(s => s.BattleObjectHooks.ViewId == damagedViewId);
            damaged.Damage(shockPower);
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

        public void SendSynchronizeStartTimeEvent()
        {
            photonView.RPC("ReceiveSynchronizeStartTimeEvent", PhotonTargets.All, UnixTime.Now);
        }

        [RPC]
        void ReceiveSynchronizeStartTimeEvent(long unixTime)
        {
            StartDateTime = UnixTime.FromUnixTime(unixTime);
        }

        public void SendPingerEvent(int submarineViewId, bool isStart)
        {
            photonView.RPC("ReceivePingerEvent", PhotonTargets.All, submarineViewId, isStart);
        }

        [RPC]
        void ReceivePingerEvent(int submarineViewId, bool isStart)
        {
            var submarine = objectContainer.Submarines
                .OfType(EnemySubmarine)
                .FirstOrDefault(s => s.BattleObjectHooks.ViewId == submarineViewId);

            if (submarine != null)
            {
                EnemyPingerCounter.Value += isStart ? 1 : -1;
            }
        }
    }
}
