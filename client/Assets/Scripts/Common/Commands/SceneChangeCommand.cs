using UnityEngine;
using UnityEngine.SceneManagement;
using Zenject;

namespace Submarine
{
    public class SceneChangeCommand : Signal<SceneNames, SceneChangeCommand>
    {
        public class Handler
        {
            public void Execute(SceneNames sceneName)
            {
                SceneManager.LoadScene(sceneName.ToString());
                Debug.Log("Loaded " + sceneName + " scene");
            }
        }
    }
}
