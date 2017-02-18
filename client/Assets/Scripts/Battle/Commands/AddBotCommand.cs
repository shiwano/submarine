using Zenject;

namespace Submarine.Battle
{
    public class AddBotCommand : Signal<AddBotCommand>
    {
        public class Handler
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
