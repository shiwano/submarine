using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine
{
    public interface ITorpedo : ITickable
    {
        TorpedoHooks Hooks { get; }
    }

    public class PlayerTorpedo : ITorpedo
    {
        public TorpedoHooks Hooks { get; private set; }

        public Vector3 Speed
        {
            get { return Hooks.transform.forward * 50f; }
        }

        public PlayerTorpedo(TorpedoHooks hooks)
        {
            Hooks = hooks;
            Hooks.OnExplode += OnExplode;
        }

        public void Tick()
        {
            Hooks.Accelerate(Speed);
        }

        private void OnExplode()
        {
            
        }
    }

    public class EnemyTorpedo : ITorpedo
    {
        public TorpedoHooks Hooks { get; private set; }

        public EnemyTorpedo(TorpedoHooks hooks)
        {
            Hooks = hooks;
        }

        public void Tick() {}
    }
}
