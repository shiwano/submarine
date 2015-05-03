using UnityEngine;
using System;

namespace Submarine
{
    public static class BattleEvent
    {
        public static Action<Photon.MonoBehaviour> OnPhotonBehaviourCreate = delegate {};
        public static Action<Photon.MonoBehaviour> OnPhotonBehaviourDestroy = delegate {};
    }
}
