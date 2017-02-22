using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomMediator : MediatorBase<RoomView>, IInitializable
    {
        [Inject]
        BattleModel battleModel;
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        StartBattleCommand startBattleCommand;
        [Inject]
        AddBotCommand addBotCommand;
        [Inject]
        RemoveBotCommand removeBotCommand;

        public void Initialize()
        {
            battleModel.OnStartAsObservable().Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            lobbyModel.JoinedRoom.Where(r => r != null).Subscribe(OnRoomChange).AddTo(view);

            view.BattleStartButtonClickedAsObservable().Subscribe(_ => OnBattleStartButtonClick()).AddTo(view);
            view.AddBotButtonClickedAsObservable().Subscribe(_ => OnAddBotButtonClick()).AddTo(view);
            view.RemoveBotButtonClickedAsObservable().Subscribe(OnRemoveBotButtonClick).AddTo(view);

            view.gameObject.SetActive(true);
        }

        void OnBattleStart()
        {
            view.gameObject.SetActive(false);
        }

        void OnRoomChange(Type.Room room)
        {
            view.RefreshRoomMembers(room.Members, room.Bots);
        }

        void OnBattleStartButtonClick()
        {
            startBattleCommand.Fire();
        }

        void OnAddBotButtonClick()
        {
            addBotCommand.Fire();
        }

        void OnRemoveBotButtonClick(Type.Bot bot)
        {
            removeBotCommand.Fire(bot.Id);
        }
    }
}
