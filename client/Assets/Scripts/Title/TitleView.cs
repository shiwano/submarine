using UnityEngine;
using UnityEngine.UI;
using UniRx;
using System;
using System.Collections;

namespace Submarine
{
    public class TitleView : MonoBehaviour
    {
        [SerializeField]
        Button startButton;

        public IObservable<Unit> StartButtonClickedAsObservable()
        {
            return startButton.OnClickAsObservable();
        }
    }
}
