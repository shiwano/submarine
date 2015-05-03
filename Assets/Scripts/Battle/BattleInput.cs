using UnityEngine;
using System;
using UniRx;
using Zenject;

namespace Submarine
{
    public class BattleInput : IInitializable
    {
        private const float mouseDraggingThreshold = 50f;
        private DateTime mouseButtonDownTime = DateTime.Now;

        public Vector3 MousePosition { get { return Input.mousePosition; } }
        public Vector3 MousePositionOnButtonDown { get; private set; }

        public float MousePressingTime
        {
            get { return IsMouseButtonPressed.Value ? (float)(DateTime.Now - mouseButtonDownTime).TotalSeconds : 0f; }
        }

        public ReactiveProperty<bool> IsMouseButtonPressed { get; private set; }

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
        }
    }
}
