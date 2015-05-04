using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine
{
    public interface ITorpedo : ITickable
    {
    }

    public class PlayerTorpedo : ITorpedo
    {
        public void Tick()
        {
            
        }
    }

    public class EnemyTorpedo : ITorpedo
    {
        public void Tick() {}
    }
}
