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
        const float sqrClickDistanceThreshold = 10f * 10f;
        DateTime touchStartTime = DateTime.Now;

        public Vector3 TouchPosition { get { return Input.mousePosition; } }
        public Vector3 TouchStartPosition { get; private set; }
        public float TouchTime { get { return IsTouched.Value ? TouchTimeInternal : 0f; } }
        float TouchTimeInternal { get { return (float)(DateTime.Now - touchStartTime).TotalSeconds; } }

        public ReactiveProperty<bool> IsTouched { get; private set; }

        public BattleInput(BattleInstaller.Settings settings)
        {
            uiSettings = settings.UI;

            IsTouched = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .Where(_ => EventSystem.current.currentSelectedGameObject == null)
                .ToReactiveProperty()
                .AddTo(disposables);

            IsTouched
                .Where(b => b)
                .Subscribe(_ =>
                {
                    touchStartTime = DateTime.Now;
                    TouchStartPosition = TouchPosition;
                })
                .AddTo(disposables);
        }

        public void Dispose()
        {
            disposables.Dispose();
        }

        public IObservable<bool> ClickedAsObservable()
        {
            return IsTouched
                .Where(b => !b &&
                    TouchTimeInternal < clickTimeThreshold &&
                    (TouchPosition - TouchStartPosition).sqrMagnitude < sqrClickDistanceThreshold
                );
        }

        public IObservable<Unit> DecoyButtonClickedAsObservable()
        {
            return uiSettings.DecoyButton.OnClickAsObservable();
        }

        public IObservable<Unit> PingerButtonClickedAsObservable()
        {
            return uiSettings.PingerButton.OnClickAsObservable();
        }

        public IObservable<Unit> LookoutButtonClickedAsObservable()
        {
            return uiSettings.LookoutButton.OnClickAsObservable();
        }
    }
}
