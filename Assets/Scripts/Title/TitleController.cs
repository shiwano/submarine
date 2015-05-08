using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class TitleController : IInitializable, IDisposable
    {
        readonly TitleInstaller.Settings sceneSettings;
        readonly MatchingService matchingService;

        public TitleController(TitleInstaller.Settings sceneSettings, MatchingService matchingService)
        {
            this.sceneSettings = sceneSettings;
            this.matchingService = matchingService;
        }

        public void Initialize()
        {
            Debug.Log("Title");
            sceneSettings.StartButton.onClick.AddListener(OnStartButtonClick);
            matchingService.onJoinRoom += OnMatchingServiceJoinRoom;
        }

        public void Dispose()
        {
            matchingService.onJoinRoom -= OnMatchingServiceJoinRoom;
        }

        void OnStartButtonClick()
        {
            Debug.Log("Click StartButton");
            sceneSettings.StartButton.onClick.RemoveListener(OnStartButtonClick);
            matchingService.JoinRoom();
        }

        void OnMatchingServiceJoinRoom()
        {
            ZenUtil.LoadScene("Battle");
        }
    }
}
