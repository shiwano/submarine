using System;
using UnityEngine;
using DG.Tweening;

namespace Submarine.Battle
{
    public abstract class ActorView : MonoBehaviour, IView, IDisposable
    {
        const float PositionY = 30f;
        const float RotationTweenDuration = 1.5f;

        public virtual void ChangeToEnemyColor() { }

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
                transform.DORotate(new Vector3(0f, 360f - value - 270f, 0f), RotationTweenDuration);
            }
        }

        public virtual void Dispose()
        {
            Destroy(gameObject);
        }
    }
}
