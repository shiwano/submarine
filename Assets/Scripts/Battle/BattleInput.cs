using UnityEngine;
using System;
using UniRx;
using Zenject;

namespace Submarine
{
    public class BattleInput : IInitializable, IDisposable
    {
        private const float mouseClickThresholdTime = 1.5f;
        private const float mouseClickThresholdDistanceSquared = 10f * 10f;
        private DateTime mouseButtonDownTime = DateTime.Now;

        public ReactiveProperty<bool> IsMouseButtonPressed { get; private set; }
        public ReactiveProperty<bool> IsMouseButtonClicked { get; private set; }

        public Vector3 MousePosition { get { return Input.mousePosition; } }
        public Vector3 MousePositionOnButtonDown { get; private set; }

        public float MousePressingTime
        {
            get { return IsMouseButtonPressed.Value ? MousePressingTimeInternal : 0f; }
        }

        private float MousePressingTimeInternal
        {
            get { return (float)(DateTime.Now - mouseButtonDownTime).TotalSeconds; }
        }

        public void Initialize()
        {
            IsMouseButtonPressed = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .ToReactiveProperty();

            IsMouseButtonPressed
                .Where(b => b)
                .Subscribe(_ =>
                {
                    mouseButtonDownTime = DateTime.Now;
                    MousePositionOnButtonDown = MousePosition;
                });

            IsMouseButtonClicked = IsMouseButtonPressed
                .Select(
                    b => !b &&
                    MousePressingTimeInternal < mouseClickThresholdTime &&
                    (MousePosition - MousePositionOnButtonDown).sqrMagnitude < mouseClickThresholdDistanceSquared)
                .ToReactiveProperty();
        }

        public void Dispose()
        {
            IsMouseButtonPressed.Dispose();
            IsMouseButtonClicked.Dispose();
        }
    }
}
