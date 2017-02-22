using UnityEngine;
using System;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class BattleInstaller : MonoInstaller
    {
        [SerializeField]
        BattleView battleView;
        [SerializeField]
        RadarView radarView;
        [SerializeField]
        RoomView roomView;
        [SerializeField]
        ResultView resultView;
        [SerializeField]
        EquipmentView equipmentView;

        public override void InstallBindings()
        {
            Container.Bind<BattleEvent.ActorCreate>().AsSingle();
            Container.Bind<BattleEvent.ActorDestroy>().AsSingle();

            Container.Bind<BattleModel>().AsSingle();
            Container.BindInterfacesAndSelfTo<BattleService>().AsSingle();

            Container.DeclareSignal<InitializeBattleCommand>().RequireHandler();
            Container.BindSignal<InitializeBattleCommand>().To<InitializeBattleCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<StartBattleCommand>().RequireHandler();
            Container.BindSignal<StartBattleCommand>().To<StartBattleCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<AddBotCommand>().RequireHandler();
            Container.BindSignal<AddBotCommand>().To<AddBotCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<RemoveBotCommand>().RequireHandler();
            Container.BindSignal<long, RemoveBotCommand>().To<RemoveBotCommand.Handler>(x => x.Execute).AsSingle();

            Container.BindMediatorAndViewAsSingle<BattleMediator, BattleView>(battleView);
            Container.BindMediatorAndViewAsSingle<RadarMediator, RadarView>(radarView);
            Container.BindMediatorAndViewAsSingle<RoomMediator, RoomView>(roomView);
            Container.BindMediatorAndViewAsSingle<ResultMediator, ResultView>(resultView);

            Container.BindInterfacesAndSelfTo<ThirdPersonCamera>().AsSingle();
            Container.BindInterfacesAndSelfTo<ActorContainer>().AsSingle();

            Container.Bind<EquipmentView>().FromInstance(equipmentView).AsSingle();
            Container.Bind<BattleInputService.IEquipmentInput>().FromInstance(equipmentView).AsSingle();
            Container.BindInterfacesAndSelfTo<BattleInputService>().AsSingle();

            Container.BindFactory<Type.Battle.Actor, bool, SubmarineFacade, SubmarineFacade.Factory>()
                .FromSubContainerResolve()
                .ByNewPrefabResource<SubmarineInstaller>(Constants.SubmarinePrefab)
                .UnderTransform(transform);

            Container.BindFactory<Type.Battle.Actor, TorpedoFacade, TorpedoFacade.Factory>()
                .FromSubContainerResolve()
                .ByNewPrefabResource<TorpedoInstaller>(Constants.TorpedoPrefab)
                .UnderTransform(transform);
        }
    }
}
