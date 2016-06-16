using System.Collections.Generic;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class ActorContainer : ITickable
    {
        [Inject]
        BattleEvent.PlayerSubmarineCreate playerSubmarineCreateEvent;
        [Inject]
        SubmarineFacade.Factory submarineFactory;
        [Inject]
        TorpedoFacade.Factory torpedoFactory;

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

        public void CreateActor(Type.Battle.Actor actor)
        {
            switch (actor.Type)
            {
                case Type.Battle.ActorType.Submarine:
                    CreateSubmarine(actor);
                    break;
                case Type.Battle.ActorType.Torpedo:
                    CreateTorpedo(actor);
                    break;
            }
        }

        public void DestroyActor(long actorId)
        {
            var actor = Get(actorId);
            if (actor != null)
            {
                actors.Remove(actorId);
                actor.Dispose();
            }
        }

        void CreateSubmarine(Type.Battle.Actor actor)
        {
            var submarine = submarineFactory.Create(actor);
            actors.Add(actor.Id, submarine);

            if (submarine.IsMine)
            {
                playerSubmarineCreateEvent.Invoke(submarine);
            }
        }

        void CreateTorpedo(Type.Battle.Actor actor)
        {
            var torpedo = torpedoFactory.Create(actor);
            actors.Add(actor.Id, torpedo);
        }
    }
}
