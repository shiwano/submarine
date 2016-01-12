using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine
{
    public class ApplicationQuitCommand : Command
    {
        public class Handler : ICommandHandler
        {
            [Inject]
            BattleService battleService;

            public void Execute()
            {
                battleService.Stop();
                Debug.Log("Game Quit");
            }
        }
    }
}
