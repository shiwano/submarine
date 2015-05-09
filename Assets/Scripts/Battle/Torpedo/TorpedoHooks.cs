using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    [RequireComponent(
        typeof(PhotonView),
        typeof(BoxCollider),
        typeof(Rigidbody)
    )]
    public class TorpedoHooks : Photon.MonoBehaviour, IBattleObjectHooks
    {
        Vector3 receivedPosition = Vector3.zero;
        Quaternion receivedRotation = Quaternion.identity;

        Rigidbody cachedRigidbody;

        public BattleObjectType Type { get { return BattleObjectType.Torpedo; } }
        public bool IsMine { get { return photonView.isMine; } }

        public event Action<int?> Striked = delegate {};

        public void Accelerate(Vector3 force)
        {
            cachedRigidbody.AddForce(force, ForceMode.Force);
        }

        public void Dispose()
        {
            if (IsMine && gameObject != null)
            {
                PhotonNetwork.Destroy(gameObject);
            }
        }

        void Awake()
        {
            cachedRigidbody = GetComponent<Rigidbody>();
            BattleEvent.BattleObjectHooksCreated(this);
        }

        void OnDestroy()
        {
            BattleEvent.BattleObjectHooksDestroyed(this);
        }

        void OnCollisionEnter(Collision collision)
        {
            if (IsMine)
            {
                var submarineHooks = collision.gameObject.GetComponent<SubmarineHooks>();
                if (submarineHooks != null && !submarineHooks.photonView.isMine)
                {
                    Striked(submarineHooks.photonView.viewID);
                }
                else
                {
                    Striked(null);
                }
            }
        }

        void Update()
        {
            if (!IsMine)
            {
                transform.position = Vector3.Lerp(transform.position, receivedPosition, Time.deltaTime * 5);
                transform.rotation = Quaternion.Lerp(transform.rotation, receivedRotation, Time.deltaTime * 5);
            }
        }

        void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info)
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
