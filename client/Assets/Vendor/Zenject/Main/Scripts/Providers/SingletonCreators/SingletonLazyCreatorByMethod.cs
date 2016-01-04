using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;

namespace Zenject
{
    internal class SingletonLazyCreatorByMethod<TConcrete> : SingletonLazyCreatorBase
    {
        readonly Func<InjectContext, TConcrete> _createMethod;

        TConcrete _instance;

        public SingletonLazyCreatorByMethod(
            SingletonId id, SingletonProviderMap owner, Func<InjectContext, TConcrete> createMethod)
            : base(id, owner)
        {
            _createMethod = createMethod;
        }

        public Func<InjectContext, TConcrete> CreateMethod
        {
            get
            {
                return _createMethod;
            }
        }

        public override object GetInstance(InjectContext context)
        {
            if (_instance == null)
            {
                _instance = _createMethod(context);

                if (_instance == null)
                {
                    throw new ZenjectResolveException(
                        "Unable to instantiate type '{0}' in SingletonLazyCreator".Fmt(context.MemberType));
                }
            }

            return _instance;
        }

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            // Can't do much here if it's a method
            yield break;
        }
    }
}

