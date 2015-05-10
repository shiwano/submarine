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
                .Where(_ => EventSystem.current.currentSelectedGameObject == null)
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

        public IObservable<bool> Clicked()
        {
            return IsTouched
                .Where(b => !b &&
                    TouchTime < clickTimeThreshold &&
                    DragAmount.sqrMagnitude < sqrClickDragLengthThreshold)
                .Select(_ => true);
        }

        public IObservable<Unit> DecoyButtonClicked()
        {
            return uiSettings.DecoyButton.OnClickAsObservable();
        }

        public IObservable<Unit> PingerButtonClicked()
        {
            return uiSettings.PingerButton.OnClickAsObservable();
        }

        public IObservable<Unit> LookoutButtonClicked()
        {
            return uiSettings.LookoutButton.OnClickAsObservable();
        }
    }
}
