using UnityEngine;
using System;
using UniRx;
using Zenject;
using UnityEngine.EventSystems;

namespace Submarine
{
    public class BattleInput : IDisposable
    {
        const float mouseClickThresholdTime = 1.5f;
        const float mouseClickThresholdDistanceSquared = 10f * 10f;
        DateTime mouseButtonDownTime = DateTime.Now;

        public ReactiveProperty<bool> IsMouseButtonPressed { get; private set; }
        public ReactiveProperty<bool> IsMouseButtonClicked { get; private set; }

        public Vector3 MousePosition { get { return Input.mousePosition; } }
        public Vector3 MousePositionOnButtonDown { get; private set; }

        public float MousePressingTime
        {
            get { return IsMouseButtonPressed.Value ? MousePressingTimeInternal : 0f; }
        }

        float MousePressingTimeInternal
        {
            get { return (float)(DateTime.Now - mouseButtonDownTime).TotalSeconds; }
        }

        public BattleInput()
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
                .Select(b => !b &&
                    EventSystem.current.currentSelectedGameObject == null &&
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
