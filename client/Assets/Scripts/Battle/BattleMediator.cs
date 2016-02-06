using UniRx;
using Zenject;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable
    {
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        BattleService battleService;
        [Inject]
        BattleView view;
        [Inject]
        StartBattleCommand startBattleCommand;

        public void Initialize()
        {
            battleService.IsStarted.Where(v => v).Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            startBattleCommand.Execute(lobbyModel.JoinedRoom.Value);
        }

        void OnBattleStart()
        {
            battleService.Api.OnPingReceiveAsObservable().Subscribe(message =>
            {
                Logger.Log(message.Message);
            });
            battleService.Api.SendPing("Hey");
        }
    }
}
