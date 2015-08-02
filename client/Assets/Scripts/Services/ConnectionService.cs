using UnityEngine;

namespace Submarine
{
    public class ConnectionService : MonoBehaviour
    {
        public bool InRoom { get { return PhotonNetwork.inRoom; } }
        public bool InLobby { get { return PhotonNetwork.insideLobby; } }
        public bool Connected { get { return PhotonNetwork.connectedAndReady; } }

        public PhotonPlayer Player { get { return PhotonNetwork.player; } }

        public bool IsMessageQueueRunning
        {
            get { return PhotonNetwork.isMessageQueueRunning; }
            set { PhotonNetwork.isMessageQueueRunning = value; }
        }

        public void Connect()
        {
            if (!Connected)
            {
                PhotonNetwork.ConnectUsingSettings(Constants.Config.Version);
            }
        }

        public void Disconnect()
        {
            if (Connected)
            {
                PhotonNetwork.Disconnect();
            }
        }
    }
}
