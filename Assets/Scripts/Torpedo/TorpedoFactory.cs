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

        public ITorpedo Create(Vector3 position, Quaternion rotation)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<TorpedoHooks>().ToMethod(c => CreateTorpedoHooks(c, position, rotation));
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

        private TorpedoHooks CreateTorpedoHooks(InjectContext context, Vector3 position, Quaternion rotation)
        {
            var go = battleService.InstantiatePhotonView(Constants.TorpedoPrefab, position, rotation);
            return go.GetComponent<TorpedoHooks>();
        }
    }
}
