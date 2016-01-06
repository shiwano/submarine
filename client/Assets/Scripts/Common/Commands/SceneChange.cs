using UnityEngine;
using System.Collections;
using Zenject;

namespace Submarine.Commands
{
    public class SceneChange
    {
        public void Execute(SceneNames sceneName)
        {
            ZenUtil.LoadScene(sceneName.ToString());
        }
    }
}
