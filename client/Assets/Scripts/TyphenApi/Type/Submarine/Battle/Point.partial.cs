using UnityEngine;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class Point
    {
        public static Point FromVector2(Vector2 v)
        {
            return new Point()
            {
                X = v.x,
                Y = v.y,
            };
        }

        public Vector2 ToVector2()
        {
            return new Vector2((float)X, (float)Y);
        }
    }
}
