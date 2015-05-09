using UnityEngine;
using System;
using UniRx;
using Zenject;
using UnityEngine.EventSystems;

namespace Submarine
{
    public class BattleInput : IDisposable
    {
        readonly CompositeDisposable disposables = new CompositeDisposable();

        const float clickTimeThreshold = 1.5f;
        const float sqrClickDistanceThreshold = 10f * 10f;
        DateTime pressStartTime = DateTime.Now;

        public Vector3 Position { get { return Input.mousePosition; } }
        public Vector3 PressStartPosition { get; private set; }
        public float PressedTime { get { return IsPressed.Value ? PressedTimeInternal : 0f; } }
        float PressedTimeInternal { get { return (float)(DateTime.Now - pressStartTime).TotalSeconds; } }

        public ReactiveProperty<bool> IsPressed { get; private set; }
        public ReactiveProperty<bool> IsClicked { get; private set; }

        public BattleInput()
        {
            IsPressed = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .ToReactiveProperty()
                .AddTo(disposables);

            IsPressed
                .Where(b => b)
                .Subscribe(_ =>
                {
                    pressStartTime = DateTime.Now;
                    PressStartPosition = Position;
                })
                .AddTo(disposables);

            IsClicked = IsPressed
                .Select(b => !b &&
                    EventSystem.current.currentSelectedGameObject == null &&
                    PressedTimeInternal < clickTimeThreshold &&
                    (Position - PressStartPosition).sqrMagnitude < sqrClickDistanceThreshold)
                .ToReactiveProperty()
                .AddTo(disposables);
        }

        public void Dispose()
        {
            disposables.Dispose();
        }
    }
}
