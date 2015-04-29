using UnityEngine;
using Zenject;

namespace Submarine
{
    public class TitleController : IInitializable
    {
        private readonly TitleInstaller.Settings sceneSettings;

        public TitleController(TitleInstaller.Settings sceneSettings)
        {
            this.sceneSettings = sceneSettings;
        }

        public void Initialize()
        {
            Debug.Log("Title");
            sceneSettings.StartButton.onClick.AddListener(OnStartButtonClick);
        }

        private void OnStartButtonClick()
        {
            Debug.Log("Click StartButton");
            ZenUtil.LoadScene("Battle");
        }
    }
}
