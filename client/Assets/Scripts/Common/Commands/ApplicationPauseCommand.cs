using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ApplicationPauseCommand : Signal<ApplicationPauseCommand>
    {
        public class Handler
        {
            public void Execute()
            {
                Debug.Log("Game Pause");
            }
        }
    }
}
