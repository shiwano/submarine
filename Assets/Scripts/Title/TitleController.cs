using UnityEngine;
using Zenject;

namespace Submarine
{
    public class TitleController : IInitializable
    {
        private readonly TitleInstaller.Settings sceneSettings;
        private readonly Transform submarine;

        public TitleController(
            TitleInstaller.Settings sceneSettings,
            [Inject("Submarine")]
            Transform submarine
        )
        {
            this.sceneSettings = sceneSettings;
            this.submarine = submarine;
        }

        public void Initialize()
        {
            Debug.Log("Title");

            sceneSettings.UI.StartButton.onClick.AddListener(OnStartButtonClick);

            submarine.localPosition = sceneSettings.Submarine.StartPosition;
            submarine.Rotate(sceneSettings.Submarine.StartRotation);
        }

        private void OnStartButtonClick()
        {
            Debug.Log("Click StartButton");
            ZenUtil.LoadScene("Battle");
        }
    }
}
