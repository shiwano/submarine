using UnityEngine;

namespace Submarine.Battle
{
    public abstract class ActorView : MonoBehaviour
    {
        const float PositionY = 20f;

        public Vector2 ActorPosition
        {
            get { return new Vector2(transform.position.x, transform.position.z); }
            set { transform.position = new Vector3(value.x, PositionY, value.y); }
        }
    }
}
