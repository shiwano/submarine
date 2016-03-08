using UnityEngine.Events;

namespace Submarine.Battle
{
    public static class BattleEvents
    {
        public class PlayerSubmarineCreate : UnityEvent<SubmarineFacade> { }
    }
}
