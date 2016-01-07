using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class SceneChange : Command<SceneNames>
    {
        public class Handler : ICommandHandler<SceneNames>
        {
            public void Execute(SceneNames sceneName)
            {
                ZenUtil.LoadScene(sceneName.ToString());
                Debug.Log("Loaded " + sceneName + " scene");
            }
        }
    }
}
