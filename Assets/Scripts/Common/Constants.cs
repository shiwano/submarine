using UnityEngine;
using System;

namespace Submarine
{
    public static class Constants
    {
        public const string Version = "0.1";
        public const int FrameRate = 30;
        public const int MapLength = 400;

        public const string SubmarinePrefab = "Battle/Submarines/Submarine";
        public const string TorpedoPrefab = "Battle/Torpedos/Torpedo";
        public const string ExplosionEffectPrefab = "Battle/Effects/Explosion";

        public static float Fps { get { return 1f / Time.deltaTime; } }
        public static float FpsRate { get { return Fps / FrameRate; } }

        public static class SceneNames
        {
            public const string Title = "Title";
            public const string Battle = "Battle";
        }
    }
}
