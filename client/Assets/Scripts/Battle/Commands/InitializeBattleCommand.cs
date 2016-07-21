using System;
using System.Linq;
using Zenject;
using Zenject.Commands;
using UniRx;

namespace Submarine.Battle
{
    public class InitializeBattleCommand : Command
    {
        public class Handler : ICommandHandler
        {
            [Inject]
            BattleService battleService;
            [Inject]
            BattleModel battleModel;
            [Inject]
            LobbyModel lobbyModel;

            public void Execute()
            {
                if (!lobbyModel.HasJoinedIntoRoom.Value)
                {
                    Logger.LogError("Not joined into a room");
                    return;
                }

                battleService.IsConnected.Where(v => v).Take(1).Subscribe(_ => OnBattleConnect());

                var room = lobbyModel.JoinedRoom.Value;
                var baseUri = new Uri(room.BattleServerBaseUri);
                var relativeUri = string.Format("rooms/{0}?room_key={1}", room.Id, room.RoomKey);
                battleService.Connect(new Uri(baseUri, relativeUri).ToString());
            }

            void OnBattleConnect()
            {
                battleModel.State.Value = BattleState.InPreparation;

                battleService.Api.OnRoomReceiveAsObservable().Subscribe(message =>
                {
                    var room = lobbyModel.JoinedRoom.Value;
                    room.Members = message.Members;
                    lobbyModel.JoinedRoom.SetValueAndForceNotify(room);
                });

                battleService.Api.OnStartReceiveAsObservable().Take(1).Subscribe(message =>
                {
                    battleModel.StartedAt = CurrentMillis.FromMilliseconds(message.StartedAt);
                    battleModel.State.Value = BattleState.InBattle;
                });

                battleService.Api.OnFinishReceiveAsObservable().Take(1).Subscribe(message =>
                {
                    battleModel.Winner = lobbyModel.JoinedRoom.Value.Members.FirstOrDefault(m => m.Id == message.WinnerUserId);
                    battleModel.FinishedAt = CurrentMillis.FromMilliseconds(message.FinishedAt);
                    lobbyModel.JoinedRoom.Value = null;
                    battleModel.State.Value = BattleState.Finish;
                });

                battleService.Api.OnNowReceiveAsObservable().Subscribe(message =>
                {
                    battleModel.Now = CurrentMillis.FromMilliseconds(message.Time);
                });
            }
        }
    }
}
