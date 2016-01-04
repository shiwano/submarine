using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;

namespace Zenject
{
    internal class SingletonLazyCreatorByInstance : SingletonLazyCreatorBase
    {
        object _instance;

        public SingletonLazyCreatorByInstance(
            SingletonId id, SingletonProviderMap owner, DiContainer container, object instance)
            : base(id, owner)
        {
            Assert.That(instance != null || container.IsValidating);
            _instance = instance;
        }

        public object Instance
        {
            get
            {
                return _instance;
            }
        }

        public override object GetInstance(InjectContext context)
        {
            return _instance;
        }

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            yield break;
        }
    }
}

