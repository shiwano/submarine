using UnityEngine;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class ApplicationPause : Command
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
