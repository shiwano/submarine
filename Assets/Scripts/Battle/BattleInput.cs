using UnityEngine;
using UniRx;
using Zenject;

namespace Submarine
{
    public class BattleInput : IInitializable
    {
        public Vector3 MousePosition { get { return Input.mousePosition; } }
        public ReactiveProperty<bool> IsMouseButtonPressed { get; private set; }

        public void Initialize()
        {
            IsMouseButtonPressed = Observable.EveryUpdate().Select(_ =>Input.GetMouseButton(0)).ToReactiveProperty();
        }
    }
}
