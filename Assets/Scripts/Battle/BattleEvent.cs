using UnityEngine;
using System;

namespace Submarine
{
    public static class BattleEvent
    {
        public static Action<IBattleObjectHooks> OnPhotonBehaviourCreate = delegate {};
        public static Action<IBattleObjectHooks> OnPhotonBehaviourDestroy = delegate {};

        public static Action<ISubmarine, ISubmarine> OnSubmarineSink = delegate {};
    }
}
