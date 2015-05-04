using UnityEngine;
using Zenject;

namespace Submarine
{
    public class TorpedoFactory
    {
        private readonly DiContainer container;
        private readonly BattleService battleService;

        public TorpedoFactory(DiContainer container, BattleService battleService)
        {
            this.container = container;
            this.battleService = battleService;
        }

        public ITorpedo Create(Vector3 position)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<TorpedoHooks>().ToMethod(c => CreateTorpedoHooks(c, position));
                return container.Instantiate<PlayerTorpedo>();
            }
        }

        public ITorpedo Create(TorpedoHooks hooks)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<TorpedoHooks>().ToInstance(hooks);
                return container.Instantiate<EnemyTorpedo>();
            }
        }

        public TorpedoHooks CreateTorpedoHooks(InjectContext context, Vector3 position)
        {
            var go = battleService.InstantiatePhotonView(Constants.TorpedoPrefab, position);
            return go.GetComponent<TorpedoHooks>();
        }
    }
}
