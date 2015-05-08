using UnityEngine;
using System.Collections;
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

        [PostInject]
        void Initialize(BattleObjectContainer objectContainer)
        {
            this.objectContainer = objectContainer;
        }
    }
}
