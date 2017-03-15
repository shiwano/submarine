using UnityEngine;
using System;
using Submarine;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class Movement
    {
        float directionForNormalizedVelocity;
        Vector2 normalizedVelocity;

        [MessagePack.IgnoreMember]
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

        [MessagePack.IgnoreMember]
        public DateTime MovedAtAsDateTime
        {
            get { return CurrentMillis.FromMilliseconds(MovedAt); }
        }
    }
}
