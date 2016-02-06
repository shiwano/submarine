using System;
using Zenject;
using Zenject.Commands;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class StartBattleCommand : Command<Type.JoinedRoom>
    {
        public class Handler : ICommandHandler<Type.JoinedRoom>
        {
            [Inject]
            BattleService battleService;

            public void Execute(Type.JoinedRoom room)
            {
                var baseUri = new Uri(room.BattleServerBaseUri);
                var relativeUri = string.Format("rooms/{0}?room_key={1}", room.Id, room.RoomKey);
                battleService.Start(new Uri(baseUri, relativeUri).ToString());
            }
        }
    }
}
