using Zenject;

namespace Submarine.Battle
{
    public class StartBattleCommand : Signal<StartBattleCommand>
    {
        public class Handler
        {
            [Inject]
            BattleService battleService;
            [Inject]
            BattleModel battleModel;

            public void Execute()
            {
                if (battleModel.State.Value != BattleState.InPreparation)
                {
                    Logger.LogError("Can't start the battle");
                    return;
                }
                battleService.Api.SendStartRequest();
            }
        }
    }
}
