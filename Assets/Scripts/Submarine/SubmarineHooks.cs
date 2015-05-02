using UnityEngine;
using System.Collections;

namespace Submarine
{
    [RequireComponent(typeof(PhotonView), typeof(Rigidbody))]
    public class SubmarineHooks : Photon.MonoBehaviour
    {
        private Vector3 receivedPosition = Vector3.zero;
        private Quaternion receivedRotation = Quaternion.identity;

        private Rigidbody cachedRigidbody;
        private const float velocityLimit = 100f;

        public bool IsMine { get { return photonView.isMine; } }

        public void AddForce(Vector3 force)
        {
            cachedRigidbody.AddForce(force, ForceMode.Force);
        }

        private void Awake()
        {
            cachedRigidbody = GetComponent<Rigidbody>();
        }

        private void Update()
        {
            if (IsMine)
            {
                cachedRigidbody.velocity = Vector3.ClampMagnitude(cachedRigidbody.velocity, velocityLimit);
            }
            else
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
