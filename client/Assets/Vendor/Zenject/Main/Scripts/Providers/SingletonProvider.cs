using System;
using System.Collections.Generic;
using ModestTree;
using System.Linq;

namespace Zenject
{
    // NOTE: we need the provider seperate from the creator because
    // if we return the same provider multiple times then the condition
    // will get over-written
    internal class SingletonProvider : ProviderBase
    {
        SingletonLazyCreatorBase _creator;

        public SingletonProvider(SingletonLazyCreatorBase creator)
        {
            _creator = creator;
        }

        public override void Dispose()
        {
            _creator.DecRefCount();
        }

        public override Type GetInstanceType()
        {
            return _creator.GetInstanceType();
        }

        public override object GetInstance(InjectContext context)
        {
            return _creator.GetInstance(context);
        }

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            return _creator.ValidateBinding(context);
        }
    }
}
