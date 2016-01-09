using UnityEngine;
using Zenject.Commands;

namespace Submarine
{
    public class ApplicationQuitCommand : Command
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
