using UnityEngine;
using Zenject;

namespace Submarine
{
    public class LookoutFactory
    {
        readonly DiContainer container;
        readonly BattleService battleService;

        public LookoutFactory(DiContainer container, BattleService battleService)
        {
            this.container = container;
            this.battleService = battleService;
        }

        public ILookout Create(Vector3 position, Quaternion rotation)
        {
            var hooks = CreateLookoutHooks(position, rotation);
            return Create(hooks);
        }

        public ILookout Create(LookoutHooks hooks)
        {
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<LookoutHooks>().ToSingleInstance(hooks);
            subContainer.Inject(hooks);

            return hooks.IsMine ?
                subContainer.Instantiate<PlayerLookout>() as ILookout :
                subContainer.Instantiate<EnemyLookout>();
        }

        LookoutHooks CreateLookoutHooks(Vector3 position, Quaternion rotation)
        {
            var go = battleService.InstantiatePhotonView(
                Constants.LookoutPrefab,
                position,
                rotation
            );
            return go.GetComponent<LookoutHooks>();
        }
    }
}
