using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class JoinIntoRoomCommand : Signal<Type.Room, JoinIntoRoomCommand>
    {
        public class Handler
        {
            [Inject]
            LobbyModel lobbyModel;
            [Inject]
            RoomService roomService;

            public void Execute(Type.Room room)
            {
                roomService.JoinIntoRoom(room).Subscribe(joinedRoom =>
                {
                    lobbyModel.JoinedRoom.Value = joinedRoom;
                });
            }
        }
    }
}
