using System.Collections.Generic;
using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class RoomListMediator : IInitializable
    {
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        CreateRoomCommand createRoomCommand;
        [Inject]
        GetRoomsCommand getRoomsCommand;
        [Inject]
        JoinIntoRoomCommand joinIntoRoomCommand;
        [Inject]
        RoomListView view;

        public void Initialize()
        {
            lobbyModel.Rooms.Subscribe(OnRoomsChange).AddTo(view);
            view.CreateRoomButtonClickedAsObservable().Subscribe(_ => OnCreateRoomButtonClick()).AddTo(view);
            view.UpdateRoomsButtonClickedAsObservable().Subscribe(_ => OnUpdateRoomsButtonClick()).AddTo(view);
        }

        void OnCreateRoomButtonClick()
        {
            createRoomCommand.Execute();
        }

        void OnUpdateRoomsButtonClick()
        {
            getRoomsCommand.Execute();
        }

        void OnRoomsChange(List<Type.Room> rooms)
        {
            view.ClearRooms();
            view.CreateRooms(rooms, OnRoomClick);
        }

        void OnRoomClick(Type.Room room)
        {
            joinIntoRoomCommand.Execute(room);
        }
    }
}
