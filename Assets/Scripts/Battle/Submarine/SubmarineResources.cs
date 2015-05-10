using UnityEngine;
using System;
using UniRx;

namespace Submarine
{
    public class SubmarineResources
    {
        const int pingerTime = 10;
        const int pingerCoolDownTime = 10;

        public IObservable<int> PingerCoolDownCounted { get; private set; }
        public ReactiveProperty<bool> IsUsingPinger { get; private set; }
        public ReactiveProperty<bool> CanUsePinger { get; private set; }

        public SubmarineResources()
        {
            IsUsingPinger = new ReactiveProperty<bool>(false);
            CanUsePinger = new ReactiveProperty<bool>(true);
        }

        public void UsePinger()
        {
            if (CanUsePinger.Value)
            {
                var observable = CreateCountdownAsObservable(pingerCoolDownTime).Publish();
                PingerCoolDownCounted = observable.AsObservable();
                PingerCoolDownCounted.Where(t => t == pingerTime).Subscribe(_ => IsUsingPinger.Value = false);
                PingerCoolDownCounted.Subscribe(_ => {}, e => {}, () => CanUsePinger.Value = true);
                observable.Connect();
                IsUsingPinger.Value = true;
                CanUsePinger.Value = false;
            }
        }

        IObservable<int> CreateCountdownAsObservable(int CountTime)
        {
            return Observable
                .Timer(TimeSpan.FromSeconds(0), TimeSpan.FromSeconds(1))
                .Select(x => (int)(CountTime - x))
                .TakeWhile(x => x > 0);
        }
    }
}
