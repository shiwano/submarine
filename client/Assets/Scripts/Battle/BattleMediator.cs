using Zenject;

namespace Submarine.Battle
{
    public class BattleMediator : IInitializable
    {
        [Inject]
        BattleView view;

        public void Initialize()
        {
        }
    }
}
