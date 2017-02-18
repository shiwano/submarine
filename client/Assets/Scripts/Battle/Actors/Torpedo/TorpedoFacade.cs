using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class TorpedoFacade : ActorFacade
    {
        public class Factory : Factory<Type.Battle.Actor, TorpedoFacade> { }

        [Inject]
        TorpedoView view;
    }
}