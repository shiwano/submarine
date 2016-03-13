using UnityEngine;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class Movement
    {
        Vector2? normalizedVelocity;
        public Vector2 NormalizedVelocity
        {
            get
            {
                if (!normalizedVelocity.HasValue)
                {
                    var directionForClient = (float)Direction;
                    normalizedVelocity = new Vector2(
                        Mathf.Cos(directionForClient * Mathf.Deg2Rad),
                        Mathf.Sin(directionForClient * Mathf.Deg2Rad)
                    );
                }
                return normalizedVelocity.Value;
            }
        }
    }
}
