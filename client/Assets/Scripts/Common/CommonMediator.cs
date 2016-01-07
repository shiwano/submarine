using Zenject;
using UniRx;

namespace Submarine
{
    public class CommonMediator : IInitializable
    {
        readonly Commands.ApplicationStart startCommand;
        readonly Commands.ApplicationPause pauseCommand;
        readonly Commands.ApplicationQuit quitCommand;

        public CommonMediator(
            Commands.ApplicationStart startCommand,
            Commands.ApplicationPause pauseCommand,
            Commands.ApplicationQuit quitCommand)
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
