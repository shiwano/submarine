using UnityEngine;
using UniRx;
using Zenject;

namespace Submarine
{
    public interface ISubmarine : ITickable, IInitializable
    {
        SubmarineHooks Hooks { get; }
    }

    public class PlayerSubmarine : ISubmarine
    {
        private readonly BattleInput input;

        public SubmarineHooks Hooks { get; private set; }

        public Vector3 Speed
        {
            get { return Hooks.transform.forward * Mathf.Lerp(0f, 50f, Mathf.Clamp01(input.MousePressingTime)); }
        }

        public Vector3 TurningEulerAngles
        {
            get { return Hooks.transform.up * (input.MousePosition.x - input.MousePositionOnButtonDown.x) * 0.01f; }
        }

        public PlayerSubmarine(SubmarineHooks hooks, BattleInput input)
        {
            Hooks = hooks;
            this.input = input;
        }

        public void Initialize()
        {
            input.IsMouseButtonPressed
                .Where(b => !b)
                .Subscribe(_ => Hooks.Brake());
            input.IsMouseButtonPressed
                .Where(b => !b)
                .Subscribe(_ => Debug.Log("Pressed"));

            input.IsMouseButtonClicked.Where(b => b).Subscribe(_ => Debug.Log("Clicked"));
        }
       
        public void Tick()
        {
            if (input.IsMouseButtonPressed.Value)
            {
                Hooks.Accelerate(Speed);
                Hooks.Turn(TurningEulerAngles);
            }
        }
    }

    public class EnemySubmarine : ISubmarine
    {
        public SubmarineHooks Hooks { get; private set; }

        public EnemySubmarine(SubmarineHooks hooks)
        {
            Hooks = hooks;
        }

        public void Initialize() {}
        public void Tick() {}
    }
}
