using UnityEngine.UI;
using System;
using UniRx;

namespace Submarine
{
    public static class UIExtensions
    {
        static TimeSpan clickThrottleDueTime = TimeSpan.FromSeconds(0.1d);

        public static IObservable<Unit> OnSingleClickAsObservable(this Button button, TimeSpan? dueTime = null)
        {
            return dueTime.HasValue ?
                button.OnClickAsObservable().ThrottleFirst(dueTime.Value) :
                button.OnClickAsObservable().ThrottleFirst(clickThrottleDueTime);
        }
    }
}
