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

        public PlayerTorpedo(TorpedoHooks hooks)
        {
            Hooks = hooks;
        }

        public void Tick()
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
