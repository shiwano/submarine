using UnityEngine;
using System;

namespace Submarine
{
    public static class BattleEvent
    {
        public static Action<IBattleObjectHooks> OnBattleObjectHooksCreate = delegate {};
        public static Action<IBattleObjectHooks> OnBattleObjectHooksDestroy = delegate {};

        public static Action<ISubmarine, ISubmarine, Vector3> OnSubmarineDamage = delegate {};
    }
}
