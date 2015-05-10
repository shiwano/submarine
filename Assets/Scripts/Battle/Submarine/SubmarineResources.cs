using System;
using UniRx;

namespace Submarine
{
    public class SubmarineResources
    {
        const int pingerCoolDownTime = 60;
        const int pingerTime = 10;

        IConnectableObservable<int> pingerCoolDownCounted;
        public IObservable<int> PingerCoolDownAsObservable { get { return pingerCoolDownCounted.AsObservable(); } }
        public ReactiveProperty<bool> CanUsePinger { get; private set; }
        public ReactiveProperty<bool> IsUsingPinger { get; private set; }

        public SubmarineResources()
        {
            IsUsingPinger = new ReactiveProperty<bool>(false);
            CanUsePinger = new ReactiveProperty<bool>(true);
        }

        public void UsePinger()
        {
            if (CanUsePinger.Value)
            {
                pingerCoolDownCounted = CreateCountdownAsObservable(pingerCoolDownTime).Publish();

                PingerCoolDownAsObservable
                    .Subscribe(_ => {}, e => {}, () => CanUsePinger.Value = true);
                PingerCoolDownAsObservable
                    .Where(t => t == pingerCoolDownTime - pingerTime)
                    .Subscribe(_ => IsUsingPinger.Value = false);

                pingerCoolDownCounted.Connect();
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
