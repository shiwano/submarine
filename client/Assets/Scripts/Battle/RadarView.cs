using UnityEngine;

namespace Submarine.Battle
{
    public class RadarView : MonoBehaviour
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
        GameObject decoyPinPrefab;
        [SerializeField]
        RectTransform pinContainer;
        [SerializeField]
        Vector2 radarSize;
    }
}
