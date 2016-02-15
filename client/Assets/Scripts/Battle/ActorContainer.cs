using System.Collections.Generic;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class ActorContainer : ITickable
    {
        [Inject]
        SubmarineFacade.Factory submarineFactory;

        List<IActor> actors = new List<IActor>();

        public void Tick()
        {
            foreach (var actor in actors)
            {
                actor.Tick();
            }
        }

        public SubmarineFacade CreateSubmarine(Type.Battle.Actor actor)
        {
            var submarine = submarineFactory.Create(actor);
            actors.Add(submarine);
            return submarine;
        }
    }
}
