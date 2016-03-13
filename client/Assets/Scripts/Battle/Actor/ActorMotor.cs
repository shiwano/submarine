using UnityEngine;
using System;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class ActorMotor
    {
        class Accelerator
        {
            public readonly bool IsAccelerating;
            public readonly double StartRate;
            public readonly double StartSpeed;
            public readonly double MaxSpeed;
            public readonly TimeSpan Duration;
            public readonly DateTime ChangedAt;

            public double Acceleration
            {
                get
                {
                    return IsAccelerating ?
                        MaxSpeed / Duration.TotalSeconds :
                        -MaxSpeed / Duration.TotalSeconds;
                }
            }

            public DateTime ReachedMaxSpeedAt
            {
                get
                {
                    var remainingRate = StartRate;
                    if (IsAccelerating)
                    {
                        remainingRate = 1 - remainingRate;
                    }
                    var remainingTime = TimeSpan.FromMilliseconds(Duration.TotalMilliseconds * remainingRate);
                    return ChangedAt + remainingTime;
                }
            }

            public Accelerator(Type.Battle.Accelerator accelerator)
            {
                IsAccelerating = accelerator.IsAccelerating;
                StartRate = accelerator.StartRate;
                MaxSpeed = accelerator.MaxSpeed;
                Duration = TimeSpan.FromMilliseconds(accelerator.Duration);
                ChangedAt = CurrentMillis.FromMilliseconds(accelerator.ChangedAt);
                StartSpeed = MaxSpeed * StartRate;
            }
        }

        [Inject]
        BattleModel battleModel;

        Type.Battle.Movement movement;
        Accelerator accelerator;
        Vector2 normalizedVelocity;

        DateTime ChangedAt
        {
            get { return CurrentMillis.FromMilliseconds(movement.MovedAt); }
        }

        public ActorMotor(Type.Battle.Actor actor)
        {
            SetMovement(actor.Movement);
        }

        public void SetMovement(Type.Battle.Movement movement)
        {
            this.movement = movement;
            accelerator = movement.Accelerator == null ?
                null :
                new Accelerator(movement.Accelerator);

            var directionForClient = (float)movement.Direction;
            normalizedVelocity = new Vector2(
                Mathf.Cos(directionForClient * Mathf.Deg2Rad),
                Mathf.Sin(directionForClient * Mathf.Deg2Rad)
            );
        }

        public Quaternion GetCurrentRotation()
        {
            return Quaternion.LookRotation(normalizedVelocity);
        }

        public Vector2 GetCurrentPosition()
        {
            if (movement == null) return Vector2.zero;
            if (accelerator == null) return movement.Position.ToVector2();

            double s1, s2 = 0d;
            if (battleModel.Now > accelerator.ReachedMaxSpeedAt)
            {
                s1 = (accelerator.ReachedMaxSpeedAt - ChangedAt).TotalSeconds;
                s2 = (battleModel.Now - accelerator.ReachedMaxSpeedAt).TotalSeconds;
            }
            else
            {
                s1 = (battleModel.Now - ChangedAt).TotalSeconds;
            }

            var p = movement.Position.ToVector2();
            var v = accelerator.StartSpeed * s1;
            var a = accelerator.Acceleration * Math.Pow(s1, 2) / 2;
            var d1 = normalizedVelocity * (float)(v + a);
            if (accelerator.IsAccelerating)
            {
                var d2 = normalizedVelocity * (float)(accelerator.MaxSpeed * s2);
                return p + d1 + d2;
            }
            return p + d1;
        }
    }
}
