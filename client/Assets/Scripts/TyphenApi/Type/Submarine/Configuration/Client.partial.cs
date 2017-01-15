using UnityEngine;

namespace TyphenApi.Type.Submarine.Configuration
{
    public partial class Client
    {
        const string Path = "Config/Config";

        public static Client Load()
        {
            var textAsset = Resources.Load<TextAsset>(Client.Path);
            Debug.Assert(textAsset != null, "Not found the Config file");
            return Client.FromJSON(textAsset.text);
        }
    }
}
