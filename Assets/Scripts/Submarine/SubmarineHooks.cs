using UnityEngine;
using System;
using DG.Tweening;

namespace Submarine
{
    [RequireComponent(
        typeof(PhotonView),
        typeof(BoxCollider),
        typeof(Rigidbody)
    )]
    public class SubmarineHooks : Photon.MonoBehaviour, IBattleObjectHooks
    {
        [SerializeField]
        private GameObject model;
        [SerializeField]
        private Transform launchSite;
        [SerializeField]
        private GameObject streamEffect;
        [SerializeField]
        private Material enemySubmarineMaterial;

        private Vector3 receivedPosition = Vector3.zero;
        private Quaternion receivedRotation = Quaternion.identity;
        private Quaternion receivedModelRotation = Quaternion.identity;

        private Rigidbody cachedRigidbody;
        private Tweener floatingTweaner;

        private const float velocityLimit = 200f;
        private const float dragOnAccelerate = 0.5f;
        private const float dragOnBrake = 2f;

        public BattleObjectType Type { get { return BattleObjectType.Submarine; } }
        public Vector3 LaunchSitePosition { get { return launchSite.position; } }

        public void Accelerate(Vector3 force)
        {
            cachedRigidbody.drag = dragOnAccelerate;
            cachedRigidbody.AddForce(force, ForceMode.Force);
        }

        public void Turn(Vector3 eulerAngles)
        {
            transform.Rotate(eulerAngles);
            model.transform.localRotation = Quaternion.Euler(new Vector3(0f, 0f, -eulerAngles.y * 15f));
        }

        public void Brake()
        {
            cachedRigidbody.drag = dragOnBrake;
            model.transform.localRotation = Quaternion.identity;
        }

        public void Sink()
        {
            floatingTweaner.Pause();
            cachedRigidbody.useGravity = true;
            streamEffect.SetActive(true);
        }

        private void Awake()
        {
            cachedRigidbody = GetComponent<Rigidbody>();
            BattleEvent.OnPhotonBehaviourCreate(this);

            if (!photonView.isMine)
            {
                model.GetComponent<MeshRenderer>().material = enemySubmarineMaterial;
            }
        }

        private void OnDestroy()
        {
            BattleEvent.OnPhotonBehaviourDestroy(this);
        }

        private void Start()
        {
            floatingTweaner = model.transform.DOLocalMoveY(-2.5f, 3f).SetEase(Ease.InOutQuad).SetLoops(-1, LoopType.Yoyo);
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
                model.transform.rotation = Quaternion.Lerp(model.transform.rotation, receivedModelRotation, Time.deltaTime * 5);
            }
        }
     
        private void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info)
        {
            if (stream.isWriting)
            {
                stream.SendNext(transform.position);
                stream.SendNext(transform.rotation);
                stream.SendNext(model.transform.rotation);
            }
            else
            {
                receivedPosition = (Vector3)stream.ReceiveNext();
                receivedRotation = (Quaternion)stream.ReceiveNext();
                receivedModelRotation = (Quaternion)stream.ReceiveNext();
            }
        }
    }
}
