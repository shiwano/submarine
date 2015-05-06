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
        private GameObject explosionEffectPrefab;

        public event Action<int> OnHitEnemySubmarine = delegate {};

        private Vector3 receivedPosition = Vector3.zero;
        private Quaternion receivedRotation = Quaternion.identity;

        private Rigidbody cachedRigidbody;
        private const float velocityLimit = 600f;

        public BattleObjectType Type { get { return BattleObjectType.Torpedo; } }

        public void Accelerate(Vector3 force)
        {
            cachedRigidbody.AddForce(force, ForceMode.Force);
        }

        public void Stop()
        {
            Destroy(gameObject);
        }

        private void Awake()
        {
            cachedRigidbody = GetComponent<Rigidbody>();
            BattleEvent.OnPhotonBehaviourCreate(this);
        }

        private void OnDestroy()
        {
            BattleEvent.OnPhotonBehaviourDestroy(this);
        }

        private void OnCollisionEnter(Collision collision)
        {
            if (!cachedRigidbody.useGravity && photonView.isMine)
            {
                var submarineHooks = collision.gameObject.GetComponent<SubmarineHooks>();
                if (submarineHooks != null && !submarineHooks.photonView.isMine)
                {
                    OnHitEnemySubmarine(submarineHooks.photonView.viewID);
                }

                photonView.RPC("Explode", PhotonTargets.All);
            }
        }

        private void Update()
        {
            if (photonView.isMine)
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

        [RPC]
        private void Explode()
        {
            var effect = Instantiate(explosionEffectPrefab);
            effect.transform.position = transform.position;

            Destroy(gameObject);
        }
    }
}
