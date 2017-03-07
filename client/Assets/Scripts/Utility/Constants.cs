using UnityEngine;

namespace Submarine
{
    public static class Constants
    {
        public const int FrameRate = 30;
        public const int MapLength = 400;

        public const string SubmarinePrefab = "Battle/Submarines/Submarine";
        public const string TorpedoPrefab = "Battle/Torpedos/Torpedo";
        public const string DecoyPrefab = "Battle/Decoys/Decoy";
        public const string WatcherPrefab = "Battle/Watchers/Watcher";
        public const string ExplosionEffectPrefab = "Battle/Effects/Explosion";

        public static float Fps { get { return 1f / Time.deltaTime; } }
        public static float FpsRate { get { return Fps / FrameRate; } }
    }
}
