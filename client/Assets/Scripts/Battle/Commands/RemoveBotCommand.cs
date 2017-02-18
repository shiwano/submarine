using Zenject;

namespace Submarine.Battle
{
    public class RemoveBotCommand : Signal<long, RemoveBotCommand>
    {
        public class Handler
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
