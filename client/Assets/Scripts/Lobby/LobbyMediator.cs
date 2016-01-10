using Zenject;

namespace Submarine.Lobby
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
