using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomMediator : IInitializable
    {
        [Inject]
        BattleModel battleModel;
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        StartBattleCommand startBattleCommand;
        [Inject]
        RoomView view;

        public void Initialize()
        {
            battleModel.OnStartAsObservable().Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            lobbyModel.JoinedRoom.Where(r => r != null).Subscribe(OnRoomChange).AddTo(view);
            view.BattleStartButtonClickedAsObservable().Subscribe(_ => OnBattleStartButtonClick()).AddTo(view);

            view.gameObject.SetActive(true);
        }

        void OnBattleStart()
        {
            view.gameObject.SetActive(false);
        }

        void OnRoomChange(Type.Room room)
        {
            view.RefreshRoomMembers(room.Members);
        }

        void OnBattleStartButtonClick()
        {
            startBattleCommand.Execute();
        }
    }
}
