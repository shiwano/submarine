using UnityEngine;

namespace Submarine.Battle
{
    [RequireComponent(typeof(WatcherView))]
    public class WatcherInstaller : ActorInstaller<WatcherFacade, WatcherView> { }
}
