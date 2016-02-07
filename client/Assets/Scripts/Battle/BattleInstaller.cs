using UnityEngine;
using System;
using Zenject;
using Zenject.Commands;
using Type = TyphenApi.Type.Submarine;

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
            Container.Bind<BattleService>().ToSingle();
            Container.Bind<BattleModel>().ToSingle();
            Container.Bind<IDisposable>().ToSingle<BattleService>();
            Container.Bind<BattleInputService>().ToSingleInstance(battleInputService);

            Container.BindCommand<InitializeBattleCommand>().HandleWithSingle<InitializeBattleCommand.Handler>();

            Container.Bind<BattleView>().ToSingleInstance(battleView);
            Container.Bind<BattleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<BattleMediator>();
            Container.Bind<ITickable>().ToSingle<BattleMediator>();

            Container.Bind<RadarView>().ToSingleInstance(radarView);
            Container.Bind<RadarMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RadarMediator>();
        }
    }
}
