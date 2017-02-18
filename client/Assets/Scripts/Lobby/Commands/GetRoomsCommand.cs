using UniRx;
using Zenject;

namespace Submarine.Lobby
{
    public class GetRoomsCommand : Signal<GetRoomsCommand>
    {
        public class Handler
        {
            [Inject]
            LobbyModel lobbyModel;
            [Inject]
            RoomService roomService;

            public void Execute()
            {
                roomService.GetRooms().Subscribe(rooms =>
                {
                    lobbyModel.Rooms.Value = rooms;
                });
            }
        }
    }
}
