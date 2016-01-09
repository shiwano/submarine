using UnityEngine;
using Zenject.Commands;

namespace Submarine
{
    public class ApplicationStartCommand : Command
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
