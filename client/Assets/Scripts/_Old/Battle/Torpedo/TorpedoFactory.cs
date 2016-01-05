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
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<TorpedoHooks>().ToSingleInstance(hooks);
            subContainer.Inject(hooks);

            return hooks.IsMine ?
                subContainer.Instantiate<PlayerTorpedo>() as ITorpedo :
                subContainer.Instantiate<EnemyTorpedo>();
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
