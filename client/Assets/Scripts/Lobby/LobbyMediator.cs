using UniRx;
using Zenject;

namespace Submarine.Lobby
{
    public class LobbyMediator : IInitializable
    {
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        BattleService battleService;
        [Inject]
        StartBattleCommand startBattleCommand;
        [Inject]
        LobbyView view;

        public void Initialize()
        {
            lobbyModel.HasJoinedIntoRoom.Where(v => v).Take(1).Subscribe(_ => OnJoinIntoRoom()).AddTo(view);
            battleService.IsStarted.Where(v => v).Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
        }

        void OnJoinIntoRoom()
        {
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
