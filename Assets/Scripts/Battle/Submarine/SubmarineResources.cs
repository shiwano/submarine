using System;
using System.Collections.Generic;
using System.Linq;
using UniRx;

namespace Submarine
{
    public class SubmarineResources
    {
        public class Resource
        {
            readonly int cooldownTime;
            readonly int usingTime;

            IConnectableObservable<int> coolDownAsConnectable;
            public IObservable<int> CoolDownAsObservable { get { return coolDownAsConnectable.AsObservable(); } }
            public ReactiveProperty<bool> CanUse { get; private set; }
            public ReactiveProperty<bool> IsUsing { get; private set; }

            public Resource(int cooldownTime, int usingTime = 0)
            {
                this.cooldownTime = cooldownTime;
                this.usingTime = usingTime;

                CanUse = new ReactiveProperty<bool>(true);
                IsUsing = new ReactiveProperty<bool>(false);
            }

            public void Use()
            {
                if (CanUse.Value)
                {
                    coolDownAsConnectable = CreateCountdownAsObservable(cooldownTime).Publish();

                    CoolDownAsObservable
                        .Subscribe(_ => {}, e => {}, () => CanUse.Value = true);
                    CoolDownAsObservable
                        .Where(t => t == cooldownTime - usingTime)
                        .Subscribe(_ => IsUsing.Value = false);

                    coolDownAsConnectable.Connect();
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

        public Resource Pinger { get; private set; }
        public List<Resource> Torpedos { get; private set; }

        public SubmarineResources()
        {
            Pinger = new Resource(60, 10);
            Torpedos = new List<Resource>()
            {
                new Resource(10),
                new Resource(10),
            };
        }
    }
}
