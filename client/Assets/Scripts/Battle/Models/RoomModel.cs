using System;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomModel
    {
        public readonly ReactiveProperty<Type.Room> Room;

        public RoomModel()
        {
            Room = new ReactiveProperty<Type.Room>();
        }
    }
}
