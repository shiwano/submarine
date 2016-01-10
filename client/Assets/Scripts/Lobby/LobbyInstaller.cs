using UnityEngine;
using Zenject;

namespace Submarine.Title
{
    public class LobbyInstaller : MonoInstaller
    {
        [SerializeField]
        LobbyView lobbyView;

        public override void InstallBindings()
        {
            Container.Bind<LobbyEvents>().ToSingle();

            Container.Bind<LobbyView>().ToSingleInstance(lobbyView);
            Container.Bind<LobbyMediator>().ToSingle();
            Container.Bind<IInitializable>().ToSingle<LobbyMediator>();
        }
    }
}
