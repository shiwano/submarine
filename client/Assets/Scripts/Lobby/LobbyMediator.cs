using UniRx;
using Zenject;

namespace Submarine.Lobby
{
    public class LobbyMediator : MediatorBase<LobbyView>, IInitializable
    {
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        SceneChangeCommand sceneChangeCommand;

        public void Initialize()
        {
            lobbyModel.HasJoinedIntoRoom.Where(v => v).Take(1).Subscribe(_ => OnJoinIntoRoom()).AddTo(view);
        }

        void OnJoinIntoRoom()
        {
            sceneChangeCommand.Fire(SceneNames.Battle);
        }
    }
}
