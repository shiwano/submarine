using UnityEngine;
using Zenject;

namespace Submarine
{
    public class TorpedoFactory
    {
        readonly DiContainer container;
        readonly BattleService battleService;

        public TorpedoFactory(DiContainer container, BattleService battleService)
        {
            this.container = container;
            this.battleService = battleService;
        }

        public ITorpedo Create(Vector3 position, Quaternion rotation)
        {
            var hooks = CreateTorpedoHooks(position, rotation);
            return Create(hooks);
        }

        public ITorpedo Create(TorpedoHooks hooks)
        {
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<TorpedoHooks>().ToSingleInstance(hooks);
                container.Inject(hooks);

                return hooks.IsMine ?
                    container.Instantiate<PlayerTorpedo>() as ITorpedo :
                    container.Instantiate<EnemyTorpedo>();
            }
        }

        TorpedoHooks CreateTorpedoHooks(Vector3 position, Quaternion rotation)
        {
            var go = battleService.InstantiatePhotonView(
                Constants.TorpedoPrefab,
                position,
                rotation
            );
            return go.GetComponent<TorpedoHooks>();
        }
    }
}
