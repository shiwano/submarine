using UniRx;
using Zenject;
using Zenject.Commands;

namespace Submarine.Lobby
{
    public class CreateRoomCommand : Command
    {
        public class Handler : ICommandHandler
        {
            [Inject]
            LobbyModel lobbyModel;
            [Inject]
            RoomService roomService;

            public void Execute()
            {
                roomService.CreateRoom().Subscribe(joinedRoom =>
                {
                    lobbyModel.JoinedRoom.Value = joinedRoom;
                });
            }
        }
    }
}
