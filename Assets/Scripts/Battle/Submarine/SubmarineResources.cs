using System;
using System.Collections.Generic;
using UniRx;

namespace Submarine
{
    public class SubmarineResources
    {
        public class Resource
        {
            readonly int cooldownTime;
            readonly int usingTime;

            public ReactiveProperty<bool> CanUse { get; private set; }
            public ReactiveProperty<bool> IsUsing { get; private set; }
            public ReactiveProperty<int> CountDown { get; private set; }

            public Resource(int cooldownTime, int usingTime = 0)
            {
                this.cooldownTime = cooldownTime;
                this.usingTime = usingTime;

                CanUse = new ReactiveProperty<bool>(true);
                IsUsing = new ReactiveProperty<bool>(false);
                CountDown = new ReactiveProperty<int>(0);
            }

            public void Use()
            {
                if (CanUse.Value)
                {
                    var observable = CreateCountdownAsObservable(cooldownTime).Publish();

                    observable.AsObservable()
                        .Subscribe(i => CountDown.Value = i, e => {}, () =>
                        {
                            CanUse.Value = true;
                            CountDown.Value = 0;
                        });
                    observable.AsObservable()
                        .Where(t => t == cooldownTime - usingTime)
                        .Subscribe(_ => IsUsing.Value = false);

                    observable.Connect();
                    CanUse.Value = false;
                    IsUsing.Value = true;
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

        public Resource Decoy { get; private set; }
        public Resource Pinger { get; private set; }
        public Resource Lookout { get; private set; }
        public List<Resource> Torpedos { get; private set; }

        public SubmarineResources()
        {
            Decoy = new Resource(150);
            Pinger = new Resource(50, 10);
            Lookout = new Resource(100);
            Torpedos = new List<Resource>()
            {
                new Resource(10),
                new Resource(10),
            };
        }
    }
}
