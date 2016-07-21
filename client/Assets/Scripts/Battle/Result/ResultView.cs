using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Battle
{
    public class ResultView : MonoBehaviour
    {
        [SerializeField]
        GameObject victoryEffect;
        [SerializeField]
        GameObject defeatEffect;
        [SerializeField]
        Button closeButton;

        public void ShowEffect(bool isVictory)
        {
            gameObject.SetActive(true);
            victoryEffect.gameObject.SetActive(isVictory);
            defeatEffect.gameObject.SetActive(!isVictory);
        }

        public IObservable<Unit> OnCloseButtonClickAsObservable()
        {
            return closeButton.OnSingleClickAsObservable();
        }

        void Start()
        {
            victoryEffect.gameObject.SetActive(false);
            defeatEffect.gameObject.SetActive(false);
            gameObject.SetActive(false);
        }
    }
}
