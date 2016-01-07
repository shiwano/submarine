using UnityEngine;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class ApplicationStart : Command
    {
        public class Handler : ICommandHandler
        {
            public void Execute()
            {
                Debug.Log("Game Start");
                Application.targetFrameRate = Constants.FrameRate;
            }
        }
    }
}
