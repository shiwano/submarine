using UnityEngine;

namespace Submarine
{
    [RequireComponent(typeof(PhotonView), typeof(Rigidbody))]
    public class SubmarineHooks : Photon.MonoBehaviour
    {
        private Vector3 receivedPosition = Vector3.zero;
        private Quaternion receivedRotation = Quaternion.identity;

        private Rigidbody cachedRigidbody;

        private const float velocityLimit = 350f;
        private const float dragOnAccelerate = 0.5f;
        private const float dragOnBrake = 1.5f;

        public bool IsMine { get { return photonView.isMine; } }

        public void Accelerate(Vector3 force)
        {
            cachedRigidbody.drag = dragOnAccelerate;
            cachedRigidbody.AddForce(force, ForceMode.Force);
        }

        public void Turn(Vector3 eulerAngles)
        {
            transform.Rotate(eulerAngles);
        }

        public void Brake()
        {
            cachedRigidbody.drag = dragOnBrake;
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
            Debug.Log("aaa");

            if (stream.isWriting)
            {
                stream.SendNext(transform.position);
                stream.SendNext(transform.rotation);
            }
            else
            {
                receivedPosition = (Vector3)stream.ReceiveNext();
                receivedRotation = (Quaternion)stream.ReceiveNext();
            }
        }
    }
}
