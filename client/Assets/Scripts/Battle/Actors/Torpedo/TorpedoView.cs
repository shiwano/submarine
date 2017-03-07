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

        public override void Dispose()
        {
            base.Dispose();
            Instantiate(explosionEffectPrefab, transform.position, transform.rotation);
        }
    }
}