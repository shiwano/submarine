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

        public Vector3 Speed
        {
            get { return hooks.transform.forward * Mathf.Lerp(0f, 50f, Mathf.Clamp01(input.MousePressingTime)); }
        }

        public Vector3 TurningEulerAngles
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
                hooks.Accelerate(Speed);
                hooks.Turn(TurningEulerAngles);
            }
            else
            {
                hooks.Brake();
            }
        }
    }
}
