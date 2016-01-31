using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine.Battle
{
    public class BattleInstaller : MonoInstaller
    {
        [SerializeField]
        BattleInputService battleInputService;
        [SerializeField]
        BattleView battleView;
        [SerializeField]
        RadarView radarView;

        public override void InstallBindings()
        {
            Container.Bind<BattleInputService>().ToSingleInstance(battleInputService);

            Container.Bind<BattleView>().ToSingleInstance(battleView);
            Container.Bind<BattleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<BattleMediator>();

            Container.Bind<RadarView>().ToSingleInstance(radarView);
            Container.Bind<RadarMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RadarMediator>();
        }
    }
}
