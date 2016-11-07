using System;
using System.Collections.Generic;
using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class ActorContainer : ITickable
    {
        [Inject]
        UserModel userModel;
        [Inject]
        SubmarineFacade.Factory submarineFactory;
        [Inject]
        TorpedoFacade.Factory torpedoFactory;

        readonly Dictionary<long, ActorFacade> actors = new Dictionary<long, ActorFacade>();

        public void Tick()
        {
            foreach (var actor in actors.Values)
            {
                actor.Tick();
            }
        }

        public bool TryGet(long actorId, out ActorFacade actorFacade)
        {
            return actors.TryGetValue(actorId, out actorFacade);
        }

        public ActorFacade CreateActor(Type.Battle.Actor actor)
        {
            switch (actor.Type)
            {
                case Type.Battle.ActorType.Submarine:
                    return CreateSubmarine(actor);
                case Type.Battle.ActorType.Torpedo:
                    return CreateTorpedo(actor);
                default:
                    throw new NotImplementedException("Unsupported actor type: " + actor.Type);
            }
        }

        public bool TryDestroyActor(long actorId, out ActorFacade actor)
        {
            if (TryGet(actorId, out actor))
            {
                actors.Remove(actorId);
                actor.Dispose();
                return true;
            }
            return false;
        }

        SubmarineFacade CreateSubmarine(Type.Battle.Actor actor)
        {
            var isPlayerSubmarine = actor.UserId == userModel.LoggedInUser.Value.Id;
            var submarine = submarineFactory.Create(actor, isPlayerSubmarine);
            actors.Add(actor.Id, submarine);
            return submarine;
        }

        TorpedoFacade CreateTorpedo(Type.Battle.Actor actor)
        {
            var torpedo = torpedoFactory.Create(actor);
            actors.Add(actor.Id, torpedo);
            return torpedo;
        }
    }
}
