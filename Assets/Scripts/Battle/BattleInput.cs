using UnityEngine;
using UniRx;
using Zenject;

namespace Submarine
{
    public class BattleInput : IInitializable
    {
        public Vector3 MousePosition { get { return Input.mousePosition; } }
        public ReactiveProperty<bool> IsClicked { get; private set; }

        public void Initialize()
        {
            IsClicked = Observable.EveryUpdate().Select(_ => Input.GetMouseButtonDown(0)).ToReactiveProperty();
        }
    }
}
