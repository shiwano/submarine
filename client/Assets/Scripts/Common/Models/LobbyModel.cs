using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine
{
    public class LobbyModel
    {
        public readonly ReactiveCollection<Type.Room> Rooms;
        public readonly ReactiveProperty<Type.JoinedRoom> JoinedRoom;

        public readonly ReadOnlyReactiveProperty<bool> HasJoinedIntoRoom;

        public LobbyModel()
        {
            Rooms = new ReactiveCollection<Type.Room>();
            JoinedRoom = new ReactiveProperty<Type.JoinedRoom>();

            HasJoinedIntoRoom = JoinedRoom.Select(r => r != null).ToReadOnlyReactiveProperty();
        }
    }
}
