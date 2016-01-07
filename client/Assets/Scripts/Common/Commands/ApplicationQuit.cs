using UnityEngine;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class ApplicationQuit : Command
    {
        public class Handler : ICommandHandler
        {
            public void Execute()
            {
                Debug.Log("Game Quit");
            }
        }
    }
}
