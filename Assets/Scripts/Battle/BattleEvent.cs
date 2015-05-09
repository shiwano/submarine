using UnityEngine;
using System;

namespace Submarine
{
    public static class BattleEvent
    {
        public static Action<IBattleObjectHooks> BattleObjectHooksCreatedViaNetwork = delegate {};
        public static Action<IBattleObjectHooks> BattleObjectHooksDestroyedViaNetwork = delegate {};

        public static Action<ISubmarine, ISubmarine, Vector3> SubmarineDamaged = delegate {};
    }
}
