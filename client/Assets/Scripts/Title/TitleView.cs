using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Title
{
    public class TitleView : MonoBehaviour, IView
    {
        [SerializeField]
        Button startButton;
        [SerializeField]
        Button deleteLoginDataButton;

        public IObservable<Unit> StartButtonClickedAsObservable()
        {
            return startButton.OnSingleClickAsObservable();
        }

        public IObservable<Unit> DeleteLoginButtonClickedAsObservable()
        {
            return deleteLoginDataButton.OnSingleClickAsObservable();
        }
    }
}
