using UnityEngine;
using System.Collections;
using System;
using Zenject;

namespace Submarine
{
    public class BattleService : Photon.MonoBehaviour
    {
        [PostInject]
        public void Initialize(MatchingService matchingService)
        {
            if (!PhotonNetwork.inRoom)
            {
                Debug.LogError("Not in room");
            }
        }
    }
}
