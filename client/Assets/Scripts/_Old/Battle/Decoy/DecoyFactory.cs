using UnityEngine;
using Zenject;

namespace Submarine
{
    public class DecoyFactory
    {
        readonly DiContainer container;
        readonly BattleService battleService;

        public DecoyFactory(DiContainer container, BattleService battleService)
        {
            this.container = container;
            this.battleService = battleService;
        }

        public IDecoy Create(Vector3 position, Quaternion rotation)
        {
            var hooks = CreateDecoyHooks(position, rotation);
            return Create(hooks);
        }

        public IDecoy Create(DecoyHooks hooks)
        {
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<DecoyHooks>().ToSingleInstance(hooks);
            subContainer.Inject(hooks);

            return hooks.IsMine ?
                subContainer.Instantiate<PlayerDecoy>() as IDecoy :
                subContainer.Instantiate<EnemyDecoy>();
        }

        DecoyHooks CreateDecoyHooks(Vector3 position, Quaternion rotation)
        {
            var go = battleService.InstantiatePhotonView(
                Constants.DecoyPrefab,
                position,
                rotation
            );
            return go.GetComponent<DecoyHooks>();
        }
    }
}
