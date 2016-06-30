using System;
using UnityEngine;
using DG.Tweening;

namespace Submarine.Battle
{
    public abstract class ActorView : MonoBehaviour, IDisposable
    {
        const float PositionY = 30f;
        const float DirectionTweenDuration = 3f;

        public abstract void ChangeToEnemyColor();

        public Vector2 ActorPosition
        {
            get { return new Vector2(transform.position.x, transform.position.z); }
            set { transform.position = new Vector3(value.x, PositionY, value.y); }
        }

        public float ActorDirection
        {
            get
            {
                var direction = (360f - transform.eulerAngles.y + 90f);
                return direction > 360f ? direction - 360f : direction;
            }
            set
            {
                transform.eulerAngles = new Vector3(0f, 360f - value - 90f, 0f);
            }
        }

        public virtual void Dispose()
        {
            Destroy(gameObject);
        }
    }
}
