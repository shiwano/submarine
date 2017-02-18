using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ApplicationStartCommand : Signal<ApplicationStartCommand>
    {
        public class Handler
        {
            public void Execute()
            {
                Debug.Log("Game Start");
                Application.targetFrameRate = Constants.FrameRate;
            }
        }
    }
}
