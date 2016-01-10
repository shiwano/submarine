using UnityEngine;
using Zenject;

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
            Container.Bind<LobbyEvents>().ToSingle();

            Container.Bind<LobbyView>().ToSingleInstance(lobbyView);
            Container.Bind<LobbyMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<LobbyMediator>();

            Container.Bind<RoomListView>().ToSingleInstance(roomListView);
            Container.Bind<RoomListMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<RoomListMediator>();
        }
    }
}
