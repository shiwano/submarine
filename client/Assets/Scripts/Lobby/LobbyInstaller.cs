using UnityEngine;
using Zenject;
using Zenject.Commands;

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

            Container.Bind<LobbyView>().ToSingleInstance(lobbyView);
            Container.Bind<LobbyMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<LobbyMediator>();

            Container.Bind<RoomListView>().ToSingleInstance(roomListView);
            Container.Bind<RoomListMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RoomListMediator>();
        }
    }
}
