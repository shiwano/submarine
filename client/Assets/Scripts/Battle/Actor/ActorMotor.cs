using UnityEngine;
using System;
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

        readonly BattleModel battleModel;

        Type.Battle.Movement movement;
        Accelerator accelerator;

        readonly TimeSpan convergenceTime = TimeSpan.FromSeconds(1f);
        DateTime convergenceFinishesAt;
        Vector2 convergenceStartPosition;
        Vector2 convergenceFinishPosition;

        DateTime ChangedAt
        {
            get { return CurrentMillis.FromMilliseconds(movement.MovedAt); }
        }

        public ActorMotor(BattleModel battleModel, Type.Battle.Actor actor)
        {
            this.battleModel = battleModel;
            SetMovement(actor.Movement);
        }

        public void SetMovement(Type.Battle.Movement newMovement)
        {
            convergenceStartPosition = GetPosition();

            movement = newMovement;
            accelerator = movement.Accelerator == null ?
                null : new Accelerator(movement.Accelerator);

            convergenceFinishesAt = battleModel.Now + convergenceTime;
            convergenceFinishPosition = GetPosition(convergenceFinishesAt);
        }

        public Quaternion GetCurrentRotation()
        {
            return movement == null ?
                Quaternion.identity :
                Quaternion.LookRotation(movement.NormalizedVelocity);
        }

        public Vector2 GetCurrentPosition()
        {
            return battleModel.Now < convergenceFinishesAt ?
                GetConvergencePosition() :
                GetPosition();
        }

        Vector2 GetConvergencePosition()
        {
            var t = (convergenceFinishesAt - battleModel.Now).TotalSeconds;
            var rate = 1f - (float)(t / convergenceTime.TotalSeconds);
            return (convergenceFinishPosition - convergenceStartPosition) * rate + convergenceStartPosition;
        }

        Vector2 GetPosition()
        {
            return GetPosition(battleModel.Now);
        }

        Vector2 GetPosition(DateTime now)
        {
            if (movement == null) return Vector2.zero;
            if (accelerator == null) return movement.Position.ToVector2();

            double t1, t2 = 0d;
            if (now > accelerator.ReachedMaxSpeedAt)
            {
                t1 = (accelerator.ReachedMaxSpeedAt - ChangedAt).TotalSeconds;
                t2 = (now - accelerator.ReachedMaxSpeedAt).TotalSeconds;
            }
            else
            {
                t1 = (now - ChangedAt).TotalSeconds;
            }

            var p = movement.Position.ToVector2();
            var v = accelerator.StartSpeed * t1;
            var a = accelerator.Acceleration * Math.Pow(t1, 2) / 2;
            var d1 = movement.NormalizedVelocity * (float)(v + a);
            if (accelerator.IsAccelerating)
            {
                var d2 = movement.NormalizedVelocity * (float)(accelerator.MaxSpeed * t2);
                return p + d1 + d2;
            }
            return p + d1;
        }
    }
}
