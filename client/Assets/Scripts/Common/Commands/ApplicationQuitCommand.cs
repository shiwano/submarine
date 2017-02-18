using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ApplicationQuitCommand : Signal<ApplicationQuitCommand>
    {
        public class Handler
        {
            public void Execute()
            {
                Debug.Log("Game Quit");
            }
        }
    }
}
