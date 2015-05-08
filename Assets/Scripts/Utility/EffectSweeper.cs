using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    public class EffectSweeper : MonoBehaviour
    {
        ParticleSystem particle;

        void Start()
        {
            particle = GetComponent<ParticleSystem>();

            Observable.Interval(TimeSpan.FromSeconds(particle.duration))
                .Take(1)
                .Subscribe(_ => Destroy(gameObject))
                .AddTo(this);
        }
    }
}
