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
        GameObject model;
        [SerializeField]
        Transform launchSite;
        [SerializeField]
        GameObject streamEffect;
        [SerializeField]
        Material enemySubmarineMaterial;

        Vector3 receivedPosition = Vector3.zero;
        Quaternion receivedRotation = Quaternion.identity;
        Quaternion receivedModelRotation = Quaternion.identity;

        Rigidbody myRigidbody;
        Tweener floatingTweaner;

        const float velocityLimit = 200f;
        const float dragOnAccelerate = 0.5f;
        const float dragOnBrake = 2f;

        public BattleObjectType Type { get { return BattleObjectType.Submarine; } }
        public Vector3 LaunchSitePosition { get { return launchSite.position; } }

        public void Accelerate(Vector3 force)
        {
            myRigidbody.drag = dragOnAccelerate;
            myRigidbody.AddForce(force, ForceMode.Force);
        }

        public void Turn(Vector3 eulerAngles)
        {
            transform.Rotate(eulerAngles);
            model.transform.localRotation = Quaternion.Euler(new Vector3(0f, 0f, -eulerAngles.y * 15f));
        }

        public void Brake()
        {
            myRigidbody.drag = dragOnBrake;
            model.transform.localRotation = Quaternion.identity;
        }

        public void Damage(Vector3 shockPower)
        {
            floatingTweaner.Pause();
            myRigidbody.useGravity = true;
            myRigidbody.constraints = RigidbodyConstraints.None;
            myRigidbody.AddForce(shockPower, ForceMode.Impulse);
            myRigidbody.AddTorque(shockPower, ForceMode.Impulse);
            streamEffect.SetActive(true);
        }

        void Awake()
        {
            myRigidbody = GetComponent<Rigidbody>();
            BattleEvent.BattleObjectHooksCreated(this);

            if (!photonView.isMine)
            {
                model.GetComponent<MeshRenderer>().material = enemySubmarineMaterial;
            }
        }

        void OnDestroy()
        {
            BattleEvent.BattleObjectHooksDestroyed(this);
        }

        void Start()
        {
            floatingTweaner = model.transform.DOLocalMoveY(-0.25f, 3f).SetEase(Ease.InOutQuad).SetLoops(-1, LoopType.Yoyo);
        }

        void Update()
        {
            if (photonView.isMine)
            {
                myRigidbody.velocity = Vector3.ClampMagnitude(myRigidbody.velocity, velocityLimit);
            }
            else
            {
                transform.position = Vector3.Lerp(transform.position, receivedPosition, Time.deltaTime * 5);
                transform.rotation = Quaternion.Lerp(transform.rotation, receivedRotation, Time.deltaTime * 5);
                model.transform.rotation = Quaternion.Lerp(model.transform.rotation, receivedModelRotation, Time.deltaTime * 5);
            }
        }
     
        void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info)
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
