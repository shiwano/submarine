using Zenject;

namespace Submarine
{
    public class CommonMediator : IInitializable
    {
        public void Initialize()
        {
            Logger.Log("Game starts");
        }
    }
}
