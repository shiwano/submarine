using System.Collections.Generic;
using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class RoomService
    {
        [Inject]
        TyphenApi.WebApi.Submarine webApi;

        public IObservable<Type.JoinedRoom> CreateRoom()
        {
            return webApi.CreateRoom().Send().Select(r => r.Data.Room);
        }

        public IObservable<List<Type.Room>> GetRooms()
        {
            return webApi.GetRooms().Send().Select(r => r.Data.Rooms);
        }

        public IObservable<Type.JoinedRoom> JoinIntoRoom(Type.Room room)
        {
            return webApi.JoinIntoRoom(room.Id).Send().Select(r => r.Data.Room);
        }
    }
}
