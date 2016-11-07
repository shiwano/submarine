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
        [SerializeField]
        RoomView roomView;
        [SerializeField]
        ResultView resultView;

        public override void InstallBindings()
        {
            Container.Bind<BattleEvent.ActorCreate>().ToSingle();
            Container.Bind<BattleEvent.ActorDestroy>().ToSingle();

            Container.Bind<BattleModel>().ToSingle();
            Container.Bind<BattleService>().ToSingle();
            Container.Bind<IDisposable>().ToSingle<BattleService>();
            Container.Bind<BattleInputService>().ToSingleInstance(battleInputService);

            Container.BindCommand<InitializeBattleCommand>().HandleWithSingle<InitializeBattleCommand.Handler>();
            Container.BindCommand<StartBattleCommand>().HandleWithSingle<StartBattleCommand.Handler>();
            Container.BindCommand<AddBotCommand>().HandleWithSingle<AddBotCommand.Handler>();
            Container.BindCommand<RemoveBotCommand, long>().HandleWithSingle<RemoveBotCommand.Handler>();

            Container.Bind<BattleView>().ToSingleInstance(battleView);
            Container.Bind<BattleMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<BattleMediator>();
            Container.Bind<ITickable>().ToSingle<BattleMediator>();

            Container.Bind<RadarView>().ToSingleInstance(radarView);
            Container.Bind<RadarMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RadarMediator>();

            Container.Bind<ThirdPersonCamera>().ToSingle();
            Container.Bind<ITickable>().ToSingle<ThirdPersonCamera>();

            Container.Bind<ActorContainer>().ToSingle();
            Container.Bind<ITickable>().ToSingle<ActorContainer>();

            Container.Bind<RoomView>().ToSingleInstance(roomView);
            Container.Bind<RoomMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RoomMediator>();

            Container.Bind<ResultView>().ToSingleInstance(resultView);
            Container.Bind<ResultMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<ResultMediator>();

            Container.BindFacadeFactory<Type.Battle.Actor, bool, SubmarineFacade, SubmarineFacade.Factory>(InstallSubmarineFacade);
            Container.BindFacadeFactory<Type.Battle.Actor, TorpedoFacade, TorpedoFacade.Factory>(InstallTorpedoFacade);
        }

        void InstallSubmarineFacade(DiContainer subContainer, Type.Battle.Actor actor, bool isPlayerSubmarine)
        {
            var submarinePrefab = Resources.Load<GameObject>(Constants.SubmarinePrefab);
            subContainer.Bind<ActorMotor>().ToSingle();
            subContainer.Bind<SubmarineView>().ToSinglePrefab(submarinePrefab);
            subContainer.Bind<ActorView>().ToSinglePrefab(submarinePrefab);
            subContainer.Bind<IDisposable>().ToSinglePrefab(submarinePrefab);
            subContainer.BindInstance(actor);

            if (isPlayerSubmarine)
            {
                subContainer.Bind<PlayerSubmarineMediator>().ToSingle();
                subContainer.Bind<IInitializable>().ToSingle<PlayerSubmarineMediator>();
                subContainer.Bind<ITickable>().ToSingle<PlayerSubmarineMediator>();
                subContainer.Bind<IDisposable>().ToSingle<PlayerSubmarineMediator>();
            }
        }

        void InstallTorpedoFacade(DiContainer subContainer, Type.Battle.Actor actor)
        {
            var torpedoPrefab = Resources.Load<GameObject>(Constants.TorpedoPrefab);
            subContainer.Bind<ActorMotor>().ToSingle();
            subContainer.Bind<TorpedoView>().ToSinglePrefab(torpedoPrefab);
            subContainer.Bind<ActorView>().ToSinglePrefab(torpedoPrefab);
            subContainer.Bind<IDisposable>().ToSinglePrefab(torpedoPrefab);
            subContainer.BindInstance(actor);
        }
    }
}
