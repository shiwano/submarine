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
        GameObject decoyPinPrefab;
        [SerializeField]
        RectTransform pinContainer;
        [SerializeField]
        Vector2 radarSize = new Vector2(250f, 250f);

        BattleObjectContainer objectContainer;

        readonly Dictionary<IBattleObject, RectTransform> pins = new Dictionary<IBattleObject, RectTransform>();

        public void SetPinger(bool isUsingPinger)
        {
            pinger.SetActive(isUsingPinger);
        }

        [PostInject]
        void Initialize(BattleObjectContainer objectContainer)
        {
            this.objectContainer = objectContainer;
            this.objectContainer.BattleObjectSpawned += OnBattleObjectSpawned;
            this.objectContainer.BattleObjectRemoved += OnBattleObjectRemoved;
        }

        void OnDestroy()
        {
            objectContainer.BattleObjectSpawned -= OnBattleObjectSpawned;
            objectContainer.BattleObjectRemoved -= OnBattleObjectRemoved;
        }

        void Update()
        {
            foreach (var pair in pins)
            {
                pair.Value.gameObject.SetActive(pair.Key.IsVisibleFromPlayer);
                pair.Value.localPosition = GetRadarPosition(pair.Key);
                pair.Value.localRotation = GetRadarRotation(pair.Key);
            }
        }

        Vector3 GetRadarPosition(IBattleObject battleObject)
        {
            var worldPosition = battleObject.Position;
            var scaleRate = radarSize.x / Constants.MapLength;
            return new Vector3(worldPosition.x * scaleRate, worldPosition.z * scaleRate, 0f);
        }

        Quaternion GetRadarRotation(IBattleObject battleObject)
        {
            var eulerAngles = new Vector3(0f, 0f, - battleObject.EulerAngles.y);
            return Quaternion.Euler(eulerAngles);
        }

        RectTransform CreatePin(IBattleObject battleObject)
        {
            GameObject pinPrefab;
            switch (battleObject.Type)
            {
                case BattleObjectType.Submarine:
                    pinPrefab = battleObject is PlayerSubmarine ?
                        playerPinPrefab :
                        enemyPinPrefab;
                    break;
                case BattleObjectType.Torpedo:
                    pinPrefab = torpedoPinPrefab;
                    break;
                case BattleObjectType.Decoy:
                    pinPrefab = decoyPinPrefab;
                    break;
                case BattleObjectType.Lookout:
                    pinPrefab = lookoutPinPrefab;
                    break;
                default:
                    return null;
            }
            return Instantiate(pinPrefab).GetComponent<RectTransform>();
        }

        void OnBattleObjectSpawned(IBattleObject battleObject)
        {
            var pin = CreatePin(battleObject);
            pin.SetParent(pinContainer);
            pin.localScale = Vector3.one; // FIXME
            pins.Add(battleObject, pin);
        }

        void OnBattleObjectRemoved(IBattleObject battleObject)
        {
            RectTransform pin;
            pins.TryGetValue(battleObject, out pin);

            if (pin != null)
            {
                Destroy(pin.gameObject);
                pins.Remove(battleObject);
            }
        }
    }
}
