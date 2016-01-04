using System;
using System.Collections.Generic;

#if !ZEN_NOT_UNITY3D
using UnityEngine;
#endif

namespace Zenject
{
    public class FacadeBinder<TFacade>
        where TFacade : IFacade
    {
        readonly Action<DiContainer> _installerFunc;
        readonly DiContainer _container;
        readonly string _identifier;

        public FacadeBinder(
            DiContainer container,
            string identifier,
            Action<DiContainer> installerFunc)
        {
            _identifier = identifier;
            _container = container;
            _installerFunc = installerFunc;
        }

        public void ToSingle()
        {
            AddValidator();
            _container.Bind<IInitializable>().ToLookup<TFacade>(_identifier);
            _container.Bind<IDisposable>().ToLookup<TFacade>(_identifier);
            _container.Bind<ITickable>().ToLookup<TFacade>(_identifier);
            _container.Bind<ILateTickable>().ToLookup<TFacade>(_identifier);
            _container.Bind<IFixedTickable>().ToLookup<TFacade>(_identifier);
            _container.Bind<TFacade>(_identifier).ToSingleMethod<TFacade>(CreateFacade);
        }

        void AddValidator()
        {
#if !ZEN_NOT_UNITY3D
            if (!Application.isPlaying)
#endif
            {
                // Unlike with facade factories, we don't really have something to be IValidatable
                // so we have to add a separate object for this in this case
                _container.Bind<IValidatable>().ToInstance(new Validator(_container, _installerFunc));
            }
        }

        TFacade CreateFacade(InjectContext ctx)
        {
            return FacadeFactory<TFacade>.CreateSubContainer(_container, _installerFunc)
                .Resolve<TFacade>();
        }

        class Validator : IValidatable
        {
            readonly DiContainer _container;
            readonly Action<DiContainer> _installerFunc;

            public Validator(DiContainer container, Action<DiContainer> installerFunc)
            {
                _container = container;
                _installerFunc = installerFunc;
            }

            public IEnumerable<ZenjectResolveException> Validate()
            {
                return FacadeFactory<TFacade>.Validate(_container, _installerFunc);
            }
        }
    }
}
