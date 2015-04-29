using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class CommonController : IInitializable, IDisposable
    {
        private readonly ConnectionService connection;

        public CommonController(ConnectionService connection)
        {
            this.connection = connection;
        }

        public void Initialize()
        {
            Debug.Log("Game Start");
        }

        public void Dispose()
        {
            connection.Disconnect();

            Debug.Log("Game Quit");
        }
    }
}
