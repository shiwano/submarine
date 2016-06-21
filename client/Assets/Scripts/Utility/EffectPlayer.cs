using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    public class EffectPlayer : MonoBehaviour
    {
        [SerializeField]
        GameObject effectPrefab;

        void Start()
        {
            var effect = Instantiate(effectPrefab);
            effect.transform.SetParent(transform, false);
            var particle = effect.GetComponent<ParticleSystem>();

            Observable.Interval(TimeSpan.FromSeconds(particle.duration))
                .Take(1)
                .Subscribe(_ => Destroy(gameObject))
                .AddTo(this);
        }
    }
}
