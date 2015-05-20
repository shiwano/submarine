using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine
{
    public class SubmarineFactory
    {
        readonly DiContainer container;
        readonly BattleService battleService;

        public SubmarineFactory(DiContainer container, BattleService battleService)
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
            using (BindScope scope = container.CreateScope())
            {
                scope.Bind<SubmarineHooks>().ToSingleInstance(hooks);
                container.Inject(hooks);

                if (hooks.IsMine)
                {
                    scope.Bind<SubmarineResources>().ToSingle();
                    return container.Instantiate<PlayerSubmarine>();
                }
                else
                {
                    return container.Instantiate<EnemySubmarine>();
                }
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
