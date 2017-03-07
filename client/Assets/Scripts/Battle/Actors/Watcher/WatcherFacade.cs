using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class WatcherFacade : ActorFacade
    {
        public class Factory : Factory<Type.Battle.Actor, WatcherFacade> { }

        [Inject]
        WatcherView view;
    }
}