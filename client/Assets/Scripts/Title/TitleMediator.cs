using UnityEngine;
using UniRx;
using Zenject;

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
