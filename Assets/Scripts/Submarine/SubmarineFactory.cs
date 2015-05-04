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

        public ISubmarine Create(Vector3 position)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<SubmarineHooks>().ToMethod(c => CreateSubmarineHooks(c, position));
                return container.Instantiate<PlayerSubmarine>();
            }
        }

        public ISubmarine Create(SubmarineHooks hooks)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<SubmarineHooks>().ToInstance(hooks);
                return container.Instantiate<EnemySubmarine>();
            }
        }

        private SubmarineHooks CreateSubmarineHooks(InjectContext context, Vector3 position)
        {
            var go = battleService.InstantiatePhotonView(Constants.SubmarinePrefab, position, Quaternion.identity);
            return go.GetComponent<SubmarineHooks>();
        }
    }
}
