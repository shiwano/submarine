using Zenject;
using UniRx;

namespace Submarine
{
    public class CommonMediator : IInitializable
    {
        readonly ApplicationStartCommand startCommand;
        readonly ApplicationPauseCommand pauseCommand;
        readonly ApplicationQuitCommand quitCommand;

        public CommonMediator(
            ApplicationStartCommand startCommand,
            ApplicationPauseCommand pauseCommand,
            ApplicationQuitCommand quitCommand)
        {
            this.startCommand = startCommand;
            this.pauseCommand = pauseCommand;
            this.quitCommand = quitCommand;
        }

        public void Initialize()
        {
            startCommand.Execute();

            Observable.OnceApplicationQuit().Subscribe(_ => quitCommand.Execute());
            Observable.EveryApplicationPause().Subscribe(_ => pauseCommand.Execute());
        }
    }
}
