using UnityEngine;
using System.Collections;


public class OnJoinInstantiate : MonoBehaviour
{
    public Transform ObjectToInstantiate;

    public bool InstantiateSceneObjects = false;

    public GameObject newObj;   // not used but to show that you get the GO as return

    public void OnJoinedRoom()
    {
        Vector3 pos = Vector3.zero;
        pos.x += PhotonNetwork.player.ID;

        if (!InstantiateSceneObjects)
        {

             newObj = PhotonNetwork.Instantiate(ObjectToInstantiate.name, pos, Quaternion.identity, 0, null);


            // anything you do with newObj locally is not reflected on the other clients.
            // you can add a script to the Prefab to do some instantiation in Awake() and you can call RPCs on newObj now.
        }
        else
        {
            newObj = PhotonNetwork.InstantiateSceneObject(ObjectToInstantiate.name, pos, Quaternion.identity, 0, null);
            //PhotonView pv = newObj.GetComponent<PhotonView>() as PhotonView;
            //Debug.Log(pv.ownerId + " " + pv.viewID);
        }
    }
}
