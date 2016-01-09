using UnityEngine.UI;
using System;
using UniRx;

namespace Submarine
{
    public static class UIExtensions
    {
        public static IObservable<Unit> onClickAsObservableWithThrottle(this Button button, TimeSpan? dueTime = null)
        {
            // TODO: Add ThrottleFirst operator after UniRX updating.
            return button.OnClickAsObservable();
        }
    }
}
