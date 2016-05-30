using Zenject;
using Zenject.Commands;

namespace Submarine.Battle
{
    public class StartBattleCommand : Command
    {
        public class Handler : ICommandHandler
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
