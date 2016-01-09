using UnityEngine;
using Zenject.Commands;

namespace Submarine
{
    public class ApplicationPauseCommand : Command
    {
        public class Handler : ICommandHandler
        {
            public void Execute()
            {
                Debug.Log("Game Pause");
            }
        }
    }
}
