using UnityEngine;

namespace TyphenApi.Type.Submarine
{
    public partial class Config
    {
        public const string Path = "Config/Config";

        public static Config Load()
        {
            var textAsset = Resources.Load<TextAsset>(Config.Path);
            Debug.Assert(textAsset != null, "Not found the Config file");
            return Config.FromJSON(textAsset.text);
        }
    }
}
