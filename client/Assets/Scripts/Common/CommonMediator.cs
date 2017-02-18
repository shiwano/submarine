using Zenject;
using UniRx;

namespace Submarine
{
    public class CommonMediator : IInitializable
    {
        [Inject]
        ApplicationStartCommand startCommand;
        [Inject]
        ApplicationPauseCommand pauseCommand;
        [Inject]
        ApplicationQuitCommand quitCommand;

        public void Initialize()
        {
            startCommand.Fire();

            Observable.OnceApplicationQuit().Subscribe(_ => quitCommand.Fire());
            Observable.EveryApplicationPause().Subscribe(_ => pauseCommand.Fire());
        }
    }
}
