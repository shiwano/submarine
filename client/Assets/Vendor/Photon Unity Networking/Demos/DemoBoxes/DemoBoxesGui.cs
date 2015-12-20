using UnityEngine;
using System.Collections;

public class DemoBoxesGui : MonoBehaviour
{
    public bool HideUI = false;

    void OnGUI()
    {
        if (HideUI)
        {
            return;
        }

        GUILayout.Label(PhotonNetwork.connectionStateDetailed.ToString());

        if (!PhotonNetwork.connected)
        {
            if (GUILayout.Button("Connect"))
            {
                PhotonNetwork.ConnectUsingSettings(null);
            }
        }
        else
        {
            if (GUILayout.Button("Disconnect"))
            {
                PhotonNetwork.Disconnect();
            }
        }

    }

}
