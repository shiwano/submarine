using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Title
{
    public class TitleView : MonoBehaviour
    {
        [SerializeField]
        Button startButton;
        [SerializeField]
        Button deleteLoginDataButton;

        public IObservable<Unit> StartButtonClickedAsObservable()
        {
            return startButton.OnClickAsObservable();
        }

        public IObservable<Unit> DeleteLoginButtonClickedAsObservable()
        {
            return deleteLoginDataButton.OnClickAsObservable();
        }
    }
}
