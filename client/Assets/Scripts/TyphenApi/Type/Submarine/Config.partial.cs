using UnityEngine;

namespace TyphenApi.Type.Submarine
{
    public partial class Config
    {
        #if UNITY_EDITOR
        public const string Path = "Config/Config.development";
        #else
        public const string Path = "Config/Config";
        #endif

        public static Config Load()
        {
            var textAsset = Resources.Load<TextAsset>(Config.Path);
            Debug.Assert(textAsset != null, "Not found a Config file");
            return Config.FromJSON(textAsset.text);
        }
    }
}
