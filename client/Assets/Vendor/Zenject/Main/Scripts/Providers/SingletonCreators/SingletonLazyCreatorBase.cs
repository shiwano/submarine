using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;

namespace Zenject
{
    internal abstract class SingletonLazyCreatorBase
    {
        readonly SingletonId _id;
        readonly SingletonProviderMap _owner;

        int _referenceCount;

        public SingletonLazyCreatorBase(
            SingletonId id, SingletonProviderMap owner)
        {
            _id = id;
            _owner = owner;
        }

        public SingletonId Id
        {
            get
            {
                return _id;
            }
        }

        public void IncRefCount()
        {
            _referenceCount += 1;
        }

        public void DecRefCount()
        {
            _referenceCount -= 1;

            if (_referenceCount <= 0)
            {
                _owner.RemoveCreator(Id);
            }
        }

        public Type GetInstanceType()
        {
            return _id.Type;
        }

        public abstract object GetInstance(InjectContext context);

        public abstract IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context);
    }
}

