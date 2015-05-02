using UnityEngine;
using System.Collections;

namespace Submarine
{
    [RequireComponent(typeof(PhotonView))]
    public class SubmarineHooks : Photon.MonoBehaviour
    {
        private Vector3 receivedPosition = Vector3.zero;
        private Quaternion receivedRotation = Quaternion.identity;

        public bool IsMine { get { return photonView.isMine; } }

        private void Update()
        {
            if (!IsMine)
            {
                transform.position = Vector3.Lerp(transform.position, receivedPosition, Time.deltaTime * 5);
                transform.rotation = Quaternion.Lerp(transform.rotation, receivedRotation, Time.deltaTime * 5);
            }
        }
     
        private void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info)
        {
            if (stream.isWriting)
            {
                stream.SendNext(transform.position);
                stream.SendNext(transform.rotation);
            }
            else
            {
                this.receivedPosition = (Vector3)stream.ReceiveNext();
                this.receivedRotation = (Quaternion)stream.ReceiveNext();
            }
        }
    }
}
