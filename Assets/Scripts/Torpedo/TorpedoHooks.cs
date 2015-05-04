using UnityEngine;

namespace Submarine
{
    [RequireComponent(
        typeof(PhotonView),
        typeof(BoxCollider),
        typeof(Rigidbody)
    )]
    public class TorpedoHooks : Photon.MonoBehaviour
    {
        private Vector3 receivedPosition = Vector3.zero;
        private Quaternion receivedRotation = Quaternion.identity;

        private Rigidbody cachedRigidbody;
        private const float velocityLimit = 600f;

        public bool IsMine { get { return photonView.isMine; } }

        private void Awake()
        {
            cachedRigidbody = GetComponent<Rigidbody>();
            BattleEvent.OnPhotonBehaviourCreate(this);
        }

        private void OnDestroy()
        {
            BattleEvent.OnPhotonBehaviourDestroy(this);
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
                receivedPosition = (Vector3)stream.ReceiveNext();
                receivedRotation = (Quaternion)stream.ReceiveNext();
            }
        }
    }
}
