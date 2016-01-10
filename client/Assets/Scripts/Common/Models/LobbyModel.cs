using UniRx;
using System.Collections.Generic;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class LobbyModel
    {
        public readonly ReactiveProperty<List<Type.Room>> Rooms;
        public readonly ReactiveProperty<Type.JoinedRoom> JoinedRoom;

        public readonly ReadOnlyReactiveProperty<bool> HasJoinedIntoRoom;

        public LobbyModel()
        {
            Rooms = new ReactiveProperty<List<Type.Room>>();
            JoinedRoom = new ReactiveProperty<Type.JoinedRoom>();

            HasJoinedIntoRoom = JoinedRoom.Select(r => r != null).ToReadOnlyReactiveProperty();
        }
    }
}
