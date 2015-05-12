using UnityEngine;
using DG.Tweening;

namespace Submarine
{
    public class SubmarineHooks : BattleObjectHooksBase
    {
        [SerializeField]
        GameObject model;
        [SerializeField]
        Transform torpedoLaunchSite;
        [SerializeField]
        Transform decoyLaunchSite;
        [SerializeField]
        Transform lookoutLaunchSite;
        [SerializeField]
        GameObject streamEffect;
        [SerializeField]
        Material enemySubmarineMaterial;

        Vector3 receivedPosition = Vector3.zero;
        Quaternion receivedRotation = Quaternion.identity;
        Quaternion receivedModelRotation = Quaternion.identity;

        Tweener floatingTweaner;
        Tweener turningBackRotationTweaner;

        const float dragOnAccelerate = 0.5f;
        const float dragOnBrake = 2f;

        public override BattleObjectType Type { get { return BattleObjectType.Submarine; } }

        public Vector3 TorpedoLaunchSitePosition { get { return torpedoLaunchSite.position; } }
        public Vector3 DecoyLaunchSitePosition { get { return decoyLaunchSite.position; } }
        public Quaternion DecoyLaunchSiteRotation { get { return decoyLaunchSite.rotation; } }
        public Vector3 LookoutLaunchSitePosition { get { return lookoutLaunchSite.position; } }

        public void Accelerate(Vector3 force)
        {
            Rigidbody.drag = dragOnAccelerate;
            Rigidbody.AddForce(force, ForceMode.Force);
        }

        public void Turn(Vector3 eulerAngles)
        {
            if (turningBackRotationTweaner != null && turningBackRotationTweaner.IsPlaying())
            {
                turningBackRotationTweaner.Kill();
            }
            transform.Rotate(eulerAngles);
            model.transform.localRotation = Quaternion.Euler(new Vector3(0f, 0f, -eulerAngles.y * 15f));
        }

        public void Brake()
        {
            turningBackRotationTweaner = model.transform.DOLocalRotate(Vector3.zero, 1f)
                .SetEase(Ease.OutExpo);
            Rigidbody.drag = dragOnBrake;
        }

        public void Damage(Vector3 shockPower)
        {
            floatingTweaner.Pause();
            Rigidbody.useGravity = true;
            Rigidbody.constraints = RigidbodyConstraints.None;
            Rigidbody.AddForce(shockPower, ForceMode.Impulse);
            Rigidbody.AddTorque(shockPower, ForceMode.Impulse);
            streamEffect.SetActive(true);
        }

        void Start()
        {
            floatingTweaner = model.transform.DOLocalMoveY(-0.25f, 3f)
                .SetEase(Ease.InOutQuad)
                .SetLoops(-1, LoopType.Yoyo);

            if (!IsMine)
            {
                model.GetComponent<MeshRenderer>().material = enemySubmarineMaterial;
            }
        }

        void Update()
        {
            if (!IsMine)
            {
                transform.position = Vector3.Lerp(transform.position, receivedPosition, Time.deltaTime * 5);
                transform.rotation = Quaternion.Lerp(transform.rotation, receivedRotation, Time.deltaTime * 5);
                model.transform.rotation = Quaternion.Lerp(model.transform.rotation, receivedModelRotation, Time.deltaTime * 5);
            }
        }
     
        protected override void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info)
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
