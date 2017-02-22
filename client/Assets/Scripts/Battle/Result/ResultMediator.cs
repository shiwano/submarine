using UniRx;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class ResultMediator : MediatorBase<ResultView>, IInitializable
    {
        [Inject]
        UserModel userModel;
        [Inject]
        BattleModel battleModel;
        [Inject]
        SceneChangeCommand sceneChangeCommand;

        public void Initialize()
        {
            battleModel.OnFinishAsObservable().Subscribe(_ => OnBattleFinish()).AddTo(view);
            view.OnCloseButtonClickAsObservable().Take(1).Subscribe(_ => OnClose()).AddTo(view);
        }

        void OnBattleFinish()
        {
            var isVictory = battleModel.Winner != null && userModel.LoggedInUser.Value.Id == battleModel.Winner.Id;
            view.ShowEffect(isVictory);
        }

        void OnClose()
        {
            sceneChangeCommand.Fire(SceneNames.Lobby);
        }
    }
}
