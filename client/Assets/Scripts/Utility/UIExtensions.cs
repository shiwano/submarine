using UnityEngine.UI;
using System;
using UniRx;

namespace Submarine
{
    public static class UIExtensions
    {
        static TimeSpan singleClickThrottleTime = TimeSpan.FromSeconds(0.1d);

        public static IObservable<Unit> OnSingleClickAsObservable(this Button button, TimeSpan? throttleTime = null)
        {
            return throttleTime.HasValue ?
                button.OnClickAsObservable().ThrottleFirst(throttleTime.Value) :
                button.OnClickAsObservable().ThrottleFirst(singleClickThrottleTime);
        }
    }
}
