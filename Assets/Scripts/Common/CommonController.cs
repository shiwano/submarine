using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public class CommonController : IInitializable, IDisposable
    {
        private readonly MatchingService matchingService;

        public CommonController(MatchingService matchingService)
        {
            this.matchingService = matchingService;
        }

        public void Initialize()
        {
            Debug.Log("Game Start");
        }

        public void Dispose()
        {
            matchingService.Disconnect();

            Debug.Log("Game Quit");
        }
    }
}
