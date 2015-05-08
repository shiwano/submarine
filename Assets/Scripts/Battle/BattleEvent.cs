using UnityEngine;
using System;

namespace Submarine
{
    public static class BattleEvent
    {
        public static Action<IBattleObjectHooks> BattleObjectHooksCreated = delegate {};
        public static Action<IBattleObjectHooks> BattleObjectHooksDestroyed = delegate {};

        public static Action<ISubmarine, ISubmarine, Vector3> SubmarineDamaged = delegate {};
    }
}
