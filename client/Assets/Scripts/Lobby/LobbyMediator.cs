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
        SceneChangeCommand sceneChangeCommand;
        [Inject]
        LobbyView view;

        public void Initialize()
        {
            lobbyModel.HasJoinedIntoRoom.Where(v => v).Take(1).Subscribe(_ => OnJoinIntoRoom()).AddTo(view);
        }

        void OnJoinIntoRoom()
        {
            startBattleCommand.Execute(lobbyModel.JoinedRoom.Value);
            sceneChangeCommand.Execute(SceneNames.Battle);
        }
    }
}
