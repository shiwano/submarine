using UnityEngine;
using System;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class Movement
    {
        float directionForNormalizedVelocity;
        Vector2 normalizedVelocity;

        public Vector2 NormalizedVelocity
        {
            get
            {
                var direction = (float)Direction;
                if (!Mathf.Approximately(direction, directionForNormalizedVelocity))
                {
                    directionForNormalizedVelocity = direction;
                    normalizedVelocity = new Vector2(
                        Mathf.Cos(direction * Mathf.Deg2Rad),
                        Mathf.Sin(direction * Mathf.Deg2Rad)
                    );
                }
                return normalizedVelocity;
            }
        }
    }
}
