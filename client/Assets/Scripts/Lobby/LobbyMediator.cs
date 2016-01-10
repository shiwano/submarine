using Zenject;

namespace Submarine.Title
{
    public class LobbyMediator : IInitializable
    {
        [Inject]
        LobbyEvents events;
        [Inject]
        LobbyView view;

        public void Initialize()
        {
        }
    }
}
