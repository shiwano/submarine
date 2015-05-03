using UnityEngine;
using Zenject;

namespace Submarine
{
    public class Submarine : IInitializable, ITickable
    {
        private readonly SubmarineHooks hooks;
        private readonly BattleInput input;

        public Transform Transform { get { return hooks.transform; } }
        public bool IsMine { get { return hooks.IsMine; } }

        public Vector3 MovingForce
        {
            get { return hooks.transform.forward * 30f; }
        }

        public Vector3 DraggedEulerAngles
        {
            get { return hooks.transform.up * (input.MousePosition.x - input.MousePositionOnButtonDown.x) * 0.01f; }
        }

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
                hooks.AddForce(MovingForce);
                hooks.transform.Rotate(DraggedEulerAngles);
            }
        }
    }
}
