using UnityEngine.Events;

namespace Submarine.Battle
{
    public static class BattleEvent
    {
        public class ActorCreate : UnityEvent<ActorFacade> { }
        public class ActorDestroy : UnityEvent<ActorFacade> { }
    }
}
