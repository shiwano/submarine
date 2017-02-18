using UniRx;
using Zenject;

namespace Submarine.Lobby
{
    public class CreateRoomCommand : Signal<CreateRoomCommand>
    {
        public class Handler
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
