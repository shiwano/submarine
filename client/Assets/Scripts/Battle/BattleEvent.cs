using UnityEngine.Events;

namespace Submarine.Battle
{
    public static class BattleEvent
    {
        public class PlayerSubmarineCreate : UnityEvent<SubmarineFacade> { }
    }
}
