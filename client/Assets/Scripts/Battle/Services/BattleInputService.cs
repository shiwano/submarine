using UnityEngine;
using UnityEngine.EventSystems;
using System;
using UniRx;

namespace Submarine.Battle
{
    public class BattleInputService : IDisposable
    {
        public interface IEquipmentInput
        {
            IObservable<Unit> OnDecoyUseAsObservable();
            IObservable<Unit> OnPingerUseAsObservable();
            IObservable<Unit> OnWatcherUseAsObservable();
        }

        readonly IEquipmentInput equipmentInput;
        readonly CompositeDisposable disposables = new CompositeDisposable();

        const float SqrDragAmountThresholdForClick = 10f * 10f;
        readonly TimeSpan touchTimeThresholdForClick = TimeSpan.FromSeconds(1.5d);
        readonly float halfScreenWidth = Screen.width / 2f;

        readonly IReadOnlyReactiveProperty<bool> isTouched;
        Vector3 touchStartPosition;
        DateTime touchStartTime;

        bool IsTouchingUI
        {
            get { return EventSystem.current != null && EventSystem.current.currentSelectedGameObject != null; }
        }

        Vector3 DragAmount
        {
            get { return isTouched.Value ? Input.mousePosition - touchStartPosition : Vector3.zero; }
        }

        public IReadOnlyReactiveProperty<bool> IsAccelerating
        {
            get { return isTouched; }
        }

        public float TurningRate
        {
            get { return Mathf.Clamp(DragAmount.x / halfScreenWidth, -1f, 1f); }
        }

        public BattleInputService(IEquipmentInput equipmentInput)
        {
            this.equipmentInput = equipmentInput;

            isTouched = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .Where(_ => !IsTouchingUI)
                .ToReactiveProperty()
                .AddTo(disposables);

            isTouched.Where(b => b)
                .Subscribe(_ => touchStartPosition = Input.mousePosition)
                .AddTo(disposables);

            isTouched.Where(b => b)
                .Subscribe(_ => touchStartTime = DateTime.Now)
                .AddTo(disposables);
        }

        public void Dispose()
        {
            disposables.Dispose();
        }

        public IObservable<Unit> OnTorpadeUseAsObservable() { return onClickAsObservable(); }
        public IObservable<Unit> OnDecoyUseAsObservable()   { return equipmentInput.OnDecoyUseAsObservable(); }
        public IObservable<Unit> OnPingerUseAsObservable()  { return equipmentInput.OnPingerUseAsObservable(); }
        public IObservable<Unit> OnWatcherUseAsObservable() { return equipmentInput.OnWatcherUseAsObservable(); }

        IObservable<Unit> onClickAsObservable()
        {
            return isTouched
                .Where(value =>
                {
                    if (value) return false;
                    var touchTime = DateTime.Now - touchStartTime;
                    var dragAmount = Input.mousePosition - touchStartPosition;
                    return touchTime < touchTimeThresholdForClick &&
                        dragAmount.sqrMagnitude < SqrDragAmountThresholdForClick;
                })
                .AsUnitObservable();
        }
    }
}
