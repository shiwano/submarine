using UniRx;
using Zenject;
using Zenject.Commands;

namespace Submarine.Lobby
{
    public class GetRoomsCommand : Command
    {
        public class Handler : ICommandHandler
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
