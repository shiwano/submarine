using UnityEngine;
using DG.Tweening;

namespace Submarine.Battle
{
    public abstract class ActorView : MonoBehaviour
    {
        const float PositionY = 30f;
        const float DirectionTweenDuration = 3f;

        public Vector2 ActorPosition
        {
            get { return new Vector2(transform.position.x, transform.position.z); }
            set { transform.position = new Vector3(value.x, PositionY, value.y); }
        }

        public float ActorDirection
        {
            get { return transform.localEulerAngles.y; }
            set { transform.DOLocalRotate(new Vector3(0f, value, 0f), DirectionTweenDuration).SetEase(Ease.Linear); }
        }
    }
}
