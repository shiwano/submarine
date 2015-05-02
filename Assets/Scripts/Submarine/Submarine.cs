using UnityEngine;

namespace Submarine
{
    public class Submarine
    {
        private readonly SubmarineHooks hooks;

        public Submarine(SubmarineHooks hooks)
        {
            this.hooks = hooks;
        }
    }
}
