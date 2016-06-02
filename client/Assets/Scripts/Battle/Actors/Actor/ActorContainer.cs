using System.Collections.Generic;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class ActorContainer : ITickable
    {
        [Inject]
        SubmarineFacade.Factory submarineFactory;

        Dictionary<long, ActorFacade> actors = new Dictionary<long, ActorFacade>();

        public void Tick()
        {
            foreach (var actor in actors.Values)
            {
                actor.Tick();
            }
        }

        public ActorFacade Get(long actorId)
        {
            ActorFacade actor;
            actors.TryGetValue(actorId, out actor);
            return actor;
        }

        public SubmarineFacade CreateSubmarine(Type.Battle.Actor actor)
        {
            var submarine = submarineFactory.Create(actor);
            actors.Add(actor.Id, submarine);
            return submarine;
        }
    }
}
