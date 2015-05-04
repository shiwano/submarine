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
        private readonly BattleObjectSpawner spawner;

        public SubmarineHooks Hooks { get; private set; }

        public Vector3 Speed
        {
            get { return Hooks.transform.forward * Mathf.Lerp(0f, 50f, Mathf.Clamp01(input.MousePressingTime)); }
        }

        public Vector3 TurningEulerAngles
        {
            get { return Hooks.transform.up * (input.MousePosition.x - input.MousePositionOnButtonDown.x) * 0.01f; }
        }

        public PlayerSubmarine(SubmarineHooks hooks, BattleInput input, BattleObjectSpawner spawner)
        {
            Hooks = hooks;
            this.input = input;
            this.spawner = spawner;
        }

        public void Initialize()
        {
            input.IsMouseButtonPressed
                .Where(b => !b)
                .Subscribe(_ => Hooks.Brake());

            input.IsMouseButtonClicked
                .Skip(1)
                .Where(b => b)
                .Subscribe(_ => SpawnTorpedo());
        }
       
        public void Tick()
        {
            if (input.IsMouseButtonPressed.Value)
            {
                Hooks.Accelerate(Speed);
                Hooks.Turn(TurningEulerAngles);
            }
        }

        private void SpawnTorpedo()
        {
            spawner.SpawnTorpedo(Hooks.transform.position, Hooks.transform.rotation);
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
