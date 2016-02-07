using UniRx;
using Zenject;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable, ITickable
    {
        [Inject]
        BattleModel battleModel;
        [Inject]
        BattleView view;
        [Inject]
        InitializeBattleCommand initializeBattleCommand;
        [Inject]
        SceneChangeCommand sceneChangeCommand;

        public void Initialize()
        {
            battleModel.OnPrepareAsObservable().Take(1).Subscribe(_ => OnBattlePrepare()).AddTo(view);
            battleModel.OnStartAsObservable().Take(1).Subscribe(_ => OnBattleStart()).AddTo(view);
            battleModel.OnFinishAsObservable().Take(1).Subscribe(_ => OnBattleFinish()).AddTo(view);

            initializeBattleCommand.Execute();
        }

        public void Tick()
        {
            if (battleModel.IsInBattle)
            {
                UpdateTimerText();
            }
        }

        void UpdateTimerText()
        {
            var elapsedTime = battleModel.Now - battleModel.StartedAt;
            view.TimerText.text = string.Format(
                "{0:00}:{1:00}",
                (int)elapsedTime.TotalMinutes,
                (int)elapsedTime.Seconds
            );
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
            sceneChangeCommand.Execute(SceneNames.Lobby);
        }
    }
}
