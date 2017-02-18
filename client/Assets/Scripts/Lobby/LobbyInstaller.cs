using UnityEngine;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class LobbyInstaller : MonoInstaller
    {
        [SerializeField]
        LobbyView lobbyView;
        [SerializeField]
        RoomListView roomListView;

        public override void InstallBindings()
        {
            Container.Bind<RoomService>().AsSingle();

            Container.DeclareSignal<CreateRoomCommand>().RequireHandler();
            Container.BindSignal<CreateRoomCommand>().To<CreateRoomCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<GetRoomsCommand>().RequireHandler();
            Container.BindSignal<GetRoomsCommand>().To<GetRoomsCommand.Handler>(x => x.Execute).AsSingle();
            Container.DeclareSignal<JoinIntoRoomCommand>().RequireHandler();
            Container.BindSignal<Type.Room, JoinIntoRoomCommand>().To<JoinIntoRoomCommand.Handler>(x => x.Execute).AsSingle();

            Container.BindInstance(lobbyView);
            Container.BindInterfacesAndSelfTo<LobbyMediator>().AsSingle();

            Container.BindInstance(roomListView);
            Container.BindInterfacesAndSelfTo<RoomListMediator>().AsSingle();
        }
    }
}
