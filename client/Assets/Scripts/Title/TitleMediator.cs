using UnityEngine;
using UniRx;
using Zenject;

namespace Submarine
{
    public class TitleMediator : IInitializable
    {
        readonly Commands.SceneChange sceneChangeCommand;
        readonly TitleView view;

        public TitleMediator(
            Commands.SceneChange sceneChangeCommand,
            TitleView view)
        {
            this.view = view;
            this.sceneChangeCommand = sceneChangeCommand;
        }

        public void Initialize()
        {
            view.StartButtonClickedAsObservable().Subscribe(_ =>
            {
                Debug.Log("a");
            });
        }

        void OnLoginSuccess()
        {
            sceneChangeCommand.Execute(SceneNames.Battle);
        }
    }
}
