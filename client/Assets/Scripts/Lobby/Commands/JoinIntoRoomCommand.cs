using UniRx;
using Zenject;
using Zenject.Commands;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class JoinIntoRoomCommand : Command<Type.Room>
    {
        public class Handler : ICommandHandler<Type.Room>
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
