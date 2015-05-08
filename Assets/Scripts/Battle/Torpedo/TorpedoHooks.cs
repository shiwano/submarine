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
        [SerializeField]
        GameObject explosionEffectPrefab;

        public event Action<int> OnHitEnemySubmarine = delegate {};

        Vector3 receivedPosition = Vector3.zero;
        Quaternion receivedRotation = Quaternion.identity;

        Rigidbody cachedRigidbody;
        bool hasExploded = false;

        public BattleObjectType Type { get { return BattleObjectType.Torpedo; } }

        public void Accelerate(Vector3 force)
        {
            cachedRigidbody.AddForce(force, ForceMode.Force);
        }

        public void Stop()
        {
            Destroy(gameObject);
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
            Collide(collision);
        }

        void OnCollisionStay(Collision collision)
        {
            Collide(collision);
        }

        void OnCollisionExit(Collision collision)
        {
            Collide(collision);
        }

        void Update()
        {
            if (!photonView.isMine)
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

        void Collide(Collision collision)
        {
            if (photonView.isMine && !hasExploded)
            {
                var submarineHooks = collision.gameObject.GetComponent<SubmarineHooks>();
                if (submarineHooks != null && !submarineHooks.photonView.isMine)
                {
                    OnHitEnemySubmarine(submarineHooks.photonView.viewID);
                }

                photonView.RPC("Explode", PhotonTargets.All);
                hasExploded = true;
            }
        }

        [RPC]
        void Explode()
        {
            var effect = Instantiate(explosionEffectPrefab);
            effect.transform.position = transform.position;

            PhotonNetwork.Destroy(gameObject);
        }
    }
}
