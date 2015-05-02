using UnityEngine;
using Zenject;
using UniRx;

namespace Submarine
{
    public class Submarine : IInitializable, ITickable
    {
        private readonly SubmarineHooks hooks;
        private readonly BattleInput input;

        private const float straightMovingForce = 30f;

        public Transform Transform { get { return hooks.transform; } }
        public bool IsMine { get { return hooks.IsMine; } }

        public Submarine(SubmarineHooks hooks, BattleInput input)
        {
            this.hooks = hooks;
            this.input = input;
        }

        public void Initialize()
        {
        }

        public void Tick()
        {
            if (!IsMine)
            {
                return;
            }

            if (input.IsMouseButtonPressed.Value)
            {
                var force = Vector3.forward * straightMovingForce;
                var draggingForce = Vector3.right * (input.MousePosition.x - input.MousePositionOnButtonDown.x);
                hooks.AddForce(force + draggingForce);
            }
        }
    }
}
