using UnityEngine;

namespace Submarine
{
    public class EffectSweeper : MonoBehaviour
    {
        ParticleSystem particle;

        void Start()
        {
            particle = GetComponent<ParticleSystem>();
        }

        void Update()
        {
            if (particle.isStopped)
            {
                Destroy(gameObject);
            }
        }
    }
}
