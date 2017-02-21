using UnityEngine;

namespace Submarine.Battle
{
    [RequireComponent(typeof(TorpedoView))]
    public class TorpedoInstaller : ActorInstaller<TorpedoFacade, TorpedoView> { }
}
