using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;

namespace Zenject
{
    //[System.Diagnostics.DebuggerStepThrough]
    public class SingletonProviderMap
    {
        readonly Dictionary<SingletonId, SingletonLazyCreatorBase> _creators = new Dictionary<SingletonId, SingletonLazyCreatorBase>();
        readonly DiContainer _container;

        public SingletonProviderMap(DiContainer container)
        {
            _container = container;
        }

        internal IEnumerable<SingletonLazyCreatorBase> Creators
        {
            get
            {
                return _creators.Values;
            }
        }

        internal void RemoveCreator(SingletonId id)
        {
            bool success = _creators.Remove(id);
            Assert.That(success);
        }

        public ProviderBase CreateProviderFromType(string identifier, Type concreteType)
        {
            return CreateProviderFromType(new SingletonId(identifier, concreteType));
        }

        public ProviderBase CreateProviderFromType(SingletonId singleId)
        {
            var creator = TryGetCreator<SingletonLazyCreatorByType>(singleId);

            if (creator == null)
            {
                creator = new SingletonLazyCreatorByType(singleId, this, _container);
                _creators.Add(singleId, creator);
            }

            return CreateProvider(creator);
        }

        public ProviderBase CreateProviderFromFactory<TContract, TFactory>(string identifier)
            where TFactory : IFactory<TContract>
        {
            var singleId = new SingletonId(identifier, typeof(TContract));
            var creator = TryGetCreator<SingletonLazyCreatorByFactory<TContract, TFactory>>(singleId);

            if (creator == null)
            {
                creator = new SingletonLazyCreatorByFactory<TContract, TFactory>(singleId, this, _container);
                _creators.Add(singleId, creator);
            }

            return CreateProvider(creator);
        }

        public ProviderBase CreateProviderFromMethod<TConcrete>(
            string identifier, Func<InjectContext, TConcrete> method)
        {
            var singleId = new SingletonId(identifier, typeof(TConcrete));
            var creator = TryGetCreator<SingletonLazyCreatorByMethod<TConcrete>>(singleId);

            if (creator == null)
            {
                creator = new SingletonLazyCreatorByMethod<TConcrete>(singleId, this, method);
                _creators.Add(singleId, creator);
            }
            else
            {
                if (!ReferenceEquals(creator.CreateMethod, method))
                {
                    throw new ZenjectBindException(
                        "Tried to bind multiple different methods for type '{0}' using ToSingleMethod".Fmt(singleId.Type.Name()));
                }
            }

            return CreateProvider(creator);
        }

        public ProviderBase CreateProviderFromInstance(
            string identifier, Type concreteType, object instance)
        {
            return CreateProviderFromInstance(
                new SingletonId(identifier, concreteType), instance);
        }

        public ProviderBase CreateProviderFromInstance(
            SingletonId singleId, object instance)
        {
            Assert.That(instance != null || _container.IsValidating);

            var creator = TryGetCreator<SingletonLazyCreatorByInstance>(singleId);

            if (creator == null)
            {
                creator = new SingletonLazyCreatorByInstance(singleId, this, _container, instance);
                _creators.Add(singleId, creator);
            }
            else
            {
                if (!ReferenceEquals(creator.Instance, instance))
                {
                    throw new ZenjectBindException(
                        "Tried to bind multiple different instances of type '{0}' using ToSingleInstance".Fmt(singleId.Type.Name()));
                }
            }

            return CreateProvider(creator);
        }

        SingletonProvider CreateProvider(SingletonLazyCreatorBase creator)
        {
            creator.IncRefCount();
            return new SingletonProvider(creator);
        }

        TCreator TryGetCreator<TCreator>(SingletonId singleId)
            where TCreator : SingletonLazyCreatorBase
        {
            SingletonLazyCreatorBase creator;

            if (_creators.TryGetValue(singleId, out creator))
            {
                if (creator.GetType() != typeof(TCreator))
                {
                    throw new ZenjectBindException(
                        "Cannot bind type '{0}' to multiple different kinds of singleton providers.  Singleton providers types: '{1}' and '{2}'"
                        .Fmt(singleId.Type, creator.GetType(), typeof(TCreator)));
                }

                return (TCreator)creator;
            }

            return null;
        }
    }
}
