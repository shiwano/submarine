using UnityEngine;
using Zenject;

namespace Submarine
{
    public class TitleController : IInitializable
    {
        private readonly TitleInstaller.Settings sceneSettings;
        private readonly Submarine submarine;

        public TitleController(TitleInstaller.Settings sceneSettings, Submarine submarine)
        {
            this.sceneSettings = sceneSettings;
            this.submarine = submarine;
        }

        public void Initialize()
        {
            Debug.Log("Title");

            sceneSettings.UI.StartButton.onClick.AddListener(OnStartButtonClick);

            submarine.SetPositionAndRotation(
                sceneSettings.Submarine.StartPosition,
                sceneSettings.Submarine.StartRotation
            );
        }

        private void OnStartButtonClick()
        {
            Debug.Log("Click StartButton");
            ZenUtil.LoadScene("Battle");
        }
    }
}
