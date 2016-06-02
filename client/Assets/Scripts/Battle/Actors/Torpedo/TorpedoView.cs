using UnityEngine;
using DG.Tweening;

namespace Submarine.Battle
{
    public class TorpedoView : ActorView
    {
        [SerializeField]
        GameObject model;
        [SerializeField]
        GameObject explosionEffectPrefab;

        public override void ChangeToEnemyColor() { }

        void OnDestroy()
        {
            Instantiate(explosionEffectPrefab, transform.position, transform.rotation);
        }
    }
}