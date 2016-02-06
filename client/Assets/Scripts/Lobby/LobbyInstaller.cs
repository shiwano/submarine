using UnityEngine;
using Zenject;
using Zenject.Commands;
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
            Container.Bind<RoomService>().ToSingle();

            Container.Bind<LobbyEvents>().ToSingle();
            Container.BindCommand<CreateRoomCommand>().HandleWithSingle<CreateRoomCommand.Handler>();
            Container.BindCommand<GetRoomsCommand>().HandleWithSingle<GetRoomsCommand.Handler>();
            Container.BindCommand<JoinIntoRoomCommand, Type.Room>().HandleWithSingle<JoinIntoRoomCommand.Handler>();

            Container.Bind<LobbyView>().ToSingleInstance(lobbyView);
            Container.Bind<LobbyMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<LobbyMediator>();

            Container.Bind<RoomListView>().ToSingleInstance(roomListView);
            Container.Bind<RoomListMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RoomListMediator>();
        }
    }
}
