using System;
using ModestTree;
#if !ZEN_NOT_UNITY3D
using UnityEngine;
#endif

namespace Zenject
{
    public abstract class BinderBase
    {
        readonly Type _contractType;
        readonly DiContainer _container;
        readonly string _bindIdentifier;

        public BinderBase(
            DiContainer container,
            Type contractType,
            string bindIdentifier)
        {
            _container = container;
            _contractType = contractType;
            _bindIdentifier = bindIdentifier;
        }

        protected Type ContractType
        {
            get
            {
                return _contractType;
            }
        }

        protected DiContainer Container
        {
            get
            {
                return _container;
            }
        }

        protected virtual BindingConditionSetter ToProvider(ProviderBase provider)
        {
            _container.RegisterProvider(
                provider, new BindingId(_contractType, _bindIdentifier));

            if (_contractType.IsValueType)
            {
                var nullableType = typeof(Nullable<>).MakeGenericType(_contractType);

                // Also bind to nullable primitives
                // this is useful so that we can have optional primitive dependencies
                _container.RegisterProvider(
                    provider, new BindingId(nullableType, _bindIdentifier));
            }

            return new BindingConditionSetter(provider);
        }
    }
}


