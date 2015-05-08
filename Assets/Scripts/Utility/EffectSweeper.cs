using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    public class EffectSweeper : MonoBehaviour
    {
        void Start()
        {
            var particle = GetComponent<ParticleSystem>();

            Observable.Interval(TimeSpan.FromSeconds(particle.duration))
                .Take(1)
                .Subscribe(_ => Destroy(gameObject))
                .AddTo(this);
        }
    }
}
