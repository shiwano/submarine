using System;
using Zenject;
using Zenject.Commands;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class StartBattleCommand : Command<Type.JoinedRoom>
    {
        public class Handler : ICommandHandler<Type.JoinedRoom>
        {
            [Inject]
            BattleService battleService;
            [Inject]
            BattleModel battleModel;

            public void Execute(Type.JoinedRoom room)
            {
                battleService.IsConnected.Where(v => v).Take(1).Subscribe(_ => OnBattleConnect());

                var baseUri = new Uri(room.BattleServerBaseUri);
                var relativeUri = string.Format("rooms/{0}?room_key={1}", room.Id, room.RoomKey);
                battleService.Connect(new Uri(baseUri, relativeUri).ToString());
            }

            void OnBattleConnect()
            {
                battleModel.State.Value = BattleState.InPreparation;
                battleService.Api.OnPingReceiveAsObservable().Subscribe(m => Logger.Log(m.Message));
                battleService.Api.SendPing("Hey");

                battleService.Api.OnStartReceiveAsObservable().Take(1).Subscribe(message =>
                {
                    battleModel.StartedAt = CurrentMillis.FromMilliseconds(message.StartedAt);
                    battleModel.State.Value = BattleState.InBattle;
                });

                battleService.Api.OnFinishReceiveAsObservable().Take(1).Subscribe(message =>
                {
                    battleModel.FinishedAt = CurrentMillis.FromMilliseconds(message.FinishedAt);
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
