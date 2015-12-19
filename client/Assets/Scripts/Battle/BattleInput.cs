using UnityEngine;
using System;
using UniRx;
using Zenject;
using UnityEngine.EventSystems;

namespace Submarine
{
    public class BattleInput : IDisposable
    {
        readonly BattleInstaller.Settings.UISettings uiSettings;
        readonly CompositeDisposable disposables = new CompositeDisposable();

        const float clickTimeThreshold = 1.5f;
        const float sqrClickDragLengthThreshold = 10f * 10f;

        public ReactiveProperty<bool> IsTouched { get; private set; }
        public ReactiveProperty<Vector3> TouchStartPosition { get; private set; }
        public ReactiveProperty<DateTime> TouchStartTime { get; private set; }

        public bool IsTouchingUI
        {
            get { return EventSystem.current != null && EventSystem.current.currentSelectedGameObject != null; }
        }

        public IObservable<bool> ClickedAsObservable
        {
            get
            {
                return IsTouched
                    .Where(b => !b &&
                        TouchTime < clickTimeThreshold &&
                        DragAmount.sqrMagnitude < sqrClickDragLengthThreshold)
                    .Select(_ => true);
            }
        }

        public IObservable<Unit> DecoyButtonClickedAsObservable
        {
            get { return uiSettings.DecoyButton.OnClickAsObservable(); }
        }

        public IObservable<Unit> PingerButtonClickedAsObservable
        {
            get { return uiSettings.PingerButton.OnClickAsObservable(); }
        }

        public IObservable<Unit> LookoutButtonClickedAsObservable
        {
            get { return uiSettings.LookoutButton.OnClickAsObservable(); }
        }

        public Vector3 DragAmount
        {
            get { return Input.mousePosition - TouchStartPosition.Value; }
        }

        public float TouchTime
        {
            get { return (float)(DateTime.Now - TouchStartTime.Value).TotalSeconds; }
        }

        public BattleInput(BattleInstaller.Settings settings)
        {
            uiSettings = settings.UI;

            IsTouched = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .Where(_ => !IsTouchingUI)
                .ToReactiveProperty()
                .AddTo(disposables);

            TouchStartPosition = IsTouched
                .Where(b => b)
                .Select(_ => Input.mousePosition)
                .ToReactiveProperty()
                .AddTo(disposables);

            TouchStartTime = IsTouched
                .Where(b => b)
                .Select(_ => DateTime.Now)
                .ToReactiveProperty()
                .AddTo(disposables);
        }

        public void Dispose()
        {
            disposables.Dispose();
        }
    }
}
