using UnityEngine;
using UniRx;
using Zenject;
using System;

namespace Submarine
{
    public class TitleMediator : IInitializable
    {
        [Inject]
        TitleView view;
        [Inject]
        Commands.SceneChange sceneChangeCommand;

        public void Initialize()
        {
            view.StartButtonClickedAsObservable().Subscribe(_ =>
                sceneChangeCommand.Execute(Constants.SceneNames.Battle));
        }
    }
}
