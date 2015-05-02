using UnityEngine;

namespace Submarine
{
    public class Submarine
    {
        private readonly SubmarineHooks hooks;

        public Transform Transform { get { return hooks.transform; } }

        public Submarine(SubmarineHooks hooks)
        {
            this.hooks = hooks;
        }
    }
}
