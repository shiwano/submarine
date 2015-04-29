using UnityEngine;

namespace Submarine
{
    public class ConnectionService : MonoBehaviour
    {
        public bool InRoom { get { return PhotonNetwork.inRoom; } }
        public bool InLobby { get { return PhotonNetwork.insideLobby; } }
        public bool Connected { get { return PhotonNetwork.connectedAndReady; } }

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
    }
}
