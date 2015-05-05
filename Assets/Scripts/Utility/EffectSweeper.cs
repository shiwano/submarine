using UnityEngine;

namespace Submarine
{
    public class EffectSweeper : MonoBehaviour
    {
        private ParticleSystem particle;

        private void Start()
        {
            particle = GetComponent<ParticleSystem>();
        }

        private void Update()
        {
            if (particle.isStopped)
            {
                Destroy(gameObject);
            }
        }
    }
}
