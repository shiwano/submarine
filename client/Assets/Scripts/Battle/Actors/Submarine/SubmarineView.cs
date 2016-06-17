using UnityEngine;
using DG.Tweening;

namespace Submarine.Battle
{
    public class SubmarineView : ActorView
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

        public Vector3 TorpedoLaunchSitePosition { get { return torpedoLaunchSite.position; } }
        public Vector3 DecoyLaunchSitePosition { get { return decoyLaunchSite.position; } }
        public Quaternion DecoyLaunchSiteRotation { get { return decoyLaunchSite.rotation; } }
        public Vector3 LookoutLaunchSitePosition { get { return lookoutLaunchSite.position; } }

        public void Turn(float rate)
        {
            var degree = transform.eulerAngles.y + rate * 1.5f;
            transform.localEulerAngles = new Vector3(0f, degree, 0f);
            model.transform.localRotation = Quaternion.Euler(new Vector3(0f, 0f, -rate * 25f));
        }

        public override void Dispose() { }

        public override void ChangeToEnemyColor()
        {
            model.GetComponent<MeshRenderer>().material = enemySubmarineMaterial;
        }

        void Start()
        {
            model.transform.DOLocalMoveY(-0.25f, 3f)
                .SetEase(Ease.InOutQuad)
                .SetLoops(-1, LoopType.Yoyo);
        }
    }
}
