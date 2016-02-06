using UniRx;
using Zenject;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable
    {
        [Inject]
        LobbyModel lobbyModel;
        [Inject]
        BattleModel battleModel;
        [Inject]
        BattleView view;
        [Inject]
        StartBattleCommand startBattleCommand;

        public void Initialize()
        {
            battleModel.OnPrepareAsObservable().Take(1).Subscribe(_ => OnBattlePrepare()).AddTo(view);
            battleModel.OnStartAsObservable().Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            battleModel.OnFinishAsObservable().Take(1).Subscribe(_ => OnBattleFinish()).AddTo(view);

            startBattleCommand.Execute(lobbyModel.JoinedRoom.Value);
        }

        void OnBattlePrepare()
        {
            Logger.Log("Battle Prepare");
        }

        void OnBattleStart()
        {
            Logger.Log("Battle Start");
        }

        void OnBattleFinish()
        {
            Logger.Log("Battle Finish");
        }
    }
}
