using UnityEngine;
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
            }

            connection.IsMessageQueueRunning = true;
        }
    }
}
