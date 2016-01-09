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
            startCommand.Execute();

            Observable.OnceApplicationQuit().Subscribe(_ => quitCommand.Execute());
            Observable.EveryApplicationPause().Subscribe(_ => pauseCommand.Execute());
        }
    }
}
