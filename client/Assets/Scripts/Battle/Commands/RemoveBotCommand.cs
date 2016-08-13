using Zenject;
using Zenject.Commands;

namespace Submarine.Battle
{
    public class RemoveBotCommand : Command<long>
    {
        public class Handler : ICommandHandler<long>
        {
            [Inject]
            BattleService battleService;
            [Inject]
            BattleModel battleModel;

            public void Execute(long botId)
            {
                if (battleModel.State.Value == BattleState.InPreparation)
                {
                    battleService.Api.SendRemoveBotRequest(botId);
                }
            }
        }
    }
}
