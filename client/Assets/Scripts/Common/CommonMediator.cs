using Zenject;
using UniRx;

namespace Submarine
{
    public class CommonMediator : IInitializable
    {
        [Inject]
        Commands.ApplicationStart startCommand;
        [Inject]
        Commands.ApplicationPause pauseCommand;
        [Inject]
        Commands.ApplicationQuit quitCommand;

        public void Initialize()
        {
            startCommand.Execute();

            Observable.OnceApplicationQuit().Subscribe(_ => quitCommand.Execute());
            Observable.EveryApplicationPause().Subscribe(_ => pauseCommand.Execute());
        }
    }
}
