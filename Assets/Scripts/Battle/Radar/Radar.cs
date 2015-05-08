using UnityEngine;
using System;
using System.Collections.Generic;
using Zenject;

namespace Submarine
{
    public class Radar : MonoBehaviour
    {
        [SerializeField]
        GameObject pinger;
        [SerializeField]
        GameObject playerPinPrefab;
        [SerializeField]
        GameObject enemyPinPrefab;
        [SerializeField]
        GameObject torpedoPinPrefab;
        [SerializeField]
        GameObject lookoutPinPrefab;
        [SerializeField]
        GameObject pinContainer;

        BattleObjectContainer objectContainer;

        Dictionary<IBattleObject, GameObject> pins = new Dictionary<IBattleObject, GameObject>();

        [PostInject]
        void Initialize(BattleObjectContainer objectContainer)
        {
            this.objectContainer = objectContainer;
        }
    }
}
