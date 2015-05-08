using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class CommonController : IInitializable
    {
        readonly ConnectionService connection;

        public CommonController(ConnectionService connection)
        {
            this.connection = connection;
        }

        public void Initialize()
        {
            Debug.Log("Game Start");
        }

        void OnApplicationQuit()
        {
            if (connection.Connected)
            {
                connection.Disconnect();
            }

            Debug.Log("Game Quit");
        }

        void OnApplicationPause(bool paused)
        {
            connection.IsMessageQueueRunning = !paused;

            if (paused)
            {
                Debug.Log("Game Pause");
            }
            else
            {
                Debug.Log("Game Resume");
            }
        }
    }
}
