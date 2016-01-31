using UnityEngine;
using UnityEngine.UI;
using UnityEngine.EventSystems;
using System;
using UniRx;

namespace Submarine.Battle
{
    public class BattleInputService : MonoBehaviour
    {
        [SerializeField]
        Button decoyButton;
        [SerializeField]
        Button pingerButton;
        [SerializeField]
        Button lookoutButton;

        const float clickTimeThreshold = 1.5f;
        const float sqrClickDragLengthThreshold = 10f * 10f;

        public ReactiveProperty<bool> IsTouched { get; private set; }
        public ReactiveProperty<Vector3> TouchStartPosition { get; private set; }
        public ReactiveProperty<DateTime> TouchStartTime { get; private set; }

        public bool IsTouchingUI
        {
            get { return EventSystem.current != null && EventSystem.current.currentSelectedGameObject != null; }
        }

        public Vector3 DragAmount
        {
            get { return Input.mousePosition - TouchStartPosition.Value; }
        }

        public float TouchTime
        {
            get { return (float)(DateTime.Now - TouchStartTime.Value).TotalSeconds; }
        }

        public IObservable<bool> ClickedAsObservable()
        {
            return IsTouched
                .Where(b => !b &&
                    TouchTime < clickTimeThreshold &&
                    DragAmount.sqrMagnitude < sqrClickDragLengthThreshold)
                .Select(_ => true);
        }

        public IObservable<Unit> DecoyButtonClickedAsObservable()
        {
            return decoyButton.OnClickAsObservable();
        }

        public IObservable<Unit> PingerButtonClickedAsObservable()
        {
            return pingerButton.OnClickAsObservable();
        }

        public IObservable<Unit> LookoutButtonClickedAsObservable()
        {
            return lookoutButton.OnClickAsObservable();
        }

        void Awake()
        {
            IsTouched = Observable.EveryUpdate()
                .Select(_ => Input.GetMouseButton(0))
                .Where(_ => !IsTouchingUI)
                .ToReactiveProperty();
            IsTouched.AddTo(this);

            TouchStartPosition = IsTouched
                .Where(b => b)
                .Select(_ => Input.mousePosition)
                .ToReactiveProperty();
            TouchStartPosition.AddTo(this);

            TouchStartTime = IsTouched
                .Where(b => b)
                .Select(_ => DateTime.Now)
                .ToReactiveProperty();
            TouchStartTime.AddTo(this);
        }
    }
}
