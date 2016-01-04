using UnityEngine;
using System.Collections;

namespace Submarine.Commands
{
    public class ApplicationStart
    {
        public void Execute()
        {
            Debug.Log("Game Start");
            Application.targetFrameRate = Constants.FrameRate;
        }
    }
}
