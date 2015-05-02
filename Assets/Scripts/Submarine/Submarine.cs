using UnityEngine;
using Zenject;
using UniRx;

namespace Submarine
{
    public class Submarine : IInitializable, ITickable
    {
        private readonly SubmarineHooks hooks;
        private readonly BattleInput input;

        public Transform Transform { get { return hooks.transform; } }

        public Submarine(SubmarineHooks hooks, BattleInput input)
        {
            this.hooks = hooks;
            this.input = input;
        }

        public void Initialize()
        {
            input.IsClicked.Where(b => b).Subscribe(_ => Debug.Log("Clicked"));
        }

        public void Tick()
        {
            Debug.Log("Tick");
        }
    }
}
