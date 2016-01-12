using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class SubmarineFactory
    {
        readonly DiContainer container;
        readonly OldBattleService battleService;

        public SubmarineFactory(DiContainer container, OldBattleService battleService)
        {
            this.container = container;
            this.battleService = battleService;
        }

        public PlayerSubmarine CreatePlayerSubmarine(Vector3 position)
        {
            var hooks = CreateSubmarineHooks(position);
            return Create(hooks) as PlayerSubmarine;
        }

        public ISubmarine Create(SubmarineHooks hooks)
        {
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<SubmarineHooks>().ToSingleInstance(hooks);
            subContainer.Inject(hooks);

            if (hooks.IsMine)
            {
                subContainer.Bind<SubmarineResources>().ToSingle();
                return subContainer.Instantiate<PlayerSubmarine>();
            }
            else
            {
                return subContainer.Instantiate<EnemySubmarine>();
            }
        }

        SubmarineHooks CreateSubmarineHooks(Vector3 position)
        {
            var go = battleService.InstantiatePhotonView(
                Constants.SubmarinePrefab,
                position,
                Quaternion.identity
            );
            return go.GetComponent<SubmarineHooks>();
        }
    }
}
