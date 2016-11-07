using System;
using UnityEngine;
using Type = TyphenApi.Type.Submarine;
using System.Collections.Generic;

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

        Dictionary<ActorFacade, RectTransform> pinsByActor = new Dictionary<ActorFacade, RectTransform>();
        Matrix4x4 matrixForRadarPosition;

        public void CreatePin(ActorFacade actorFacade)
        {
            GameObject pinPrefab;
            switch (actorFacade.Actor.Type)
            {
                case Type.Battle.ActorType.Submarine:
                    pinPrefab = actorFacade.IsMine ?
                        playerPinPrefab :
                        enemyPinPrefab;
                    break;
                case Type.Battle.ActorType.Torpedo:
                    pinPrefab = torpedoPinPrefab;
                    break;
                default:
                    throw new NotImplementedException("Unsupported actor type: " + actorFacade.Actor.Type);
            }
            var pin = Instantiate(pinPrefab).GetComponent<RectTransform>();
            pin.SetParent(pinContainer, true);
            pinsByActor.Add(actorFacade, pin);
        }

        public void DestroyPin(ActorFacade actorFacade)
        {
            RectTransform pin;
            if (pinsByActor.TryGetValue(actorFacade, out pin))
            {
                Destroy(pin.gameObject);
                pinsByActor.Remove(actorFacade);
            }
        }

        void Awake()
        {
            matrixForRadarPosition = Matrix4x4.zero;
            matrixForRadarPosition.m00 = radarSize.x / Constants.MapLength;;
            matrixForRadarPosition.m11 = radarSize.x / Constants.MapLength;;
        }

        void Update()
        {
            foreach (var pair in pinsByActor)
            {
                pair.Value.gameObject.SetActive(pair.Key.Actor.IsVisible);
                pair.Value.localPosition = GetRadarPosition(pair.Key);
                pair.Value.localRotation = GetRadarRotation(pair.Key);
            }
        }

        Vector3 GetRadarPosition(ActorFacade actorFacade)
        {
            var position = actorFacade.View.ActorPosition;
            return matrixForRadarPosition.MultiplyPoint3x4(new Vector3(position.x, position.y));
        }

        Quaternion GetRadarRotation(ActorFacade actorFacade)
        {
            var eulerAngles = new Vector3(0f, 0f, actorFacade.View.ActorDirection - 90f);
            return Quaternion.Euler(eulerAngles);
        }
    }
}
