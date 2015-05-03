using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class SubmarineFactory
    {
        private readonly DiContainer container;
        private readonly BattleService battleService;

        public SubmarineFactory(DiContainer container, BattleService battleService)
        {
            this.container = container;
            this.battleService = battleService;
        }

        public Submarine Create(Vector3 position)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<SubmarineHooks>().ToMethod(c => GetOrCreateSubmarineHooks(c, position));
                return container.Instantiate<Submarine>();
            }
        }

        public SubmarineHooks GetOrCreateSubmarineHooks(InjectContext context, Vector3 position)
        {
            var go = battleService.InstantiatePhotonView(Constants.SubmarinePrefab, position);
            return go.GetComponent<SubmarineHooks>();
        }
    }
}
