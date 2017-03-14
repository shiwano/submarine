using Zenject;

namespace Submarine
{
    public class CommonMediator : IInitializable
    {
        public void Initialize()
        {
            Logger.Log("Game starts");
            var config = TyphenApi.Type.Submarine.Configuration.Client.Load();
            var a = config.ToMessagePack();
            var b = TyphenApi.Type.Submarine.Configuration.Client.FromMessagePack(a);
            Logger.Log(b.ToJSON());
        }
    }
}
