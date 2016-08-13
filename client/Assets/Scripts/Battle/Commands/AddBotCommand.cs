using Zenject;
using Zenject.Commands;

namespace Submarine.Battle
{
    public class AddBotCommand : Command
    {
        public class Handler : ICommandHandler
        {
            [Inject]
            BattleService battleService;
            [Inject]
            BattleModel battleModel;

            public void Execute()
            {
                if (battleModel.State.Value == BattleState.InPreparation)
                {
                    battleService.Api.SendAddBotRequest();
                }
            }
        }
    }
}
