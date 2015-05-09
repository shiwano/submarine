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

        public BattleInput()
        {
            IsPressed = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .Where(_ => EventSystem.current.currentSelectedGameObject == null)
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
        }

        public void Dispose()
        {
            disposables.Dispose();
        }

        public IObservable<bool> ClickedAsObservable()
        {
            return IsPressed
                .Where(b => !b &&
                    PressedTimeInternal < clickTimeThreshold &&
                    (Position - PressStartPosition).sqrMagnitude < sqrClickDistanceThreshold
                );
        }
    }
}
