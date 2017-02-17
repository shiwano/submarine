using System;

#if !NOT_UNITY3D
using UnityEngine;
#endif

namespace Zenject
{
    // We wrap FromBinderGeneric because some things in there that don't really work for signals
    // like CopyInSubContainers, conditions, subcontainers
    public class SignalFromBinder<TContract> : ScopeArgNonLazyBinder
    {
        readonly BindInfo _info;
        readonly FromBinderGeneric<TContract> _subBinder;

        public SignalFromBinder(
            BindInfo info, FromBinderGeneric<TContract> subBinder)
            : base(info)
        {
            _info = info;
            _subBinder = subBinder;
        }

        public ScopeArgNonLazyBinder FromFactory<TFactory>()
            where TFactory : IFactory<TContract>
        {
            _subBinder.FromFactory<TFactory>();
            return this;
        }

        public ScopeArgNonLazyBinder FromFactory<TConcrete, TFactory>()
            where TFactory : IFactory<TConcrete>
            where TConcrete : TContract
        {
            _subBinder.FromFactory<TConcrete, TFactory>();
            return this;
        }

        public ScopeArgNonLazyBinder FromMethod(Func<InjectContext, TContract> method)
        {
            _subBinder.FromMethod(method);
            return this;
        }

        public ScopeNonLazyBinder FromResolveGetter<TObj>(Func<TObj, TContract> method)
        {
            _subBinder.FromResolveGetter<TObj>(method);
            return new ScopeNonLazyBinder(_info);
        }

        public ScopeNonLazyBinder FromResolveGetter<TObj>(object identifier, Func<TObj, TContract> method)
        {
            _subBinder.FromResolveGetter<TObj>(identifier, method);
            return new ScopeNonLazyBinder(_info);
        }

        public ScopeNonLazyBinder FromInstance(TContract instance)
        {
            _subBinder.FromInstance(instance);
            return new ScopeNonLazyBinder(_info);
        }

        public ScopeNonLazyBinder FromInstance(TContract instance, bool allowNull)
        {
            _subBinder.FromInstance(instance, allowNull);
            return new ScopeNonLazyBinder(_info);
        }

        // This is the default if nothing else is called
        public ScopeArgNonLazyBinder FromNew()
        {
            _subBinder.FromNew();
            return this;
        }

        public ScopeNonLazyBinder FromResolve()
        {
            _subBinder.FromResolve();
            return new ScopeNonLazyBinder(_info);
        }

        public ScopeNonLazyBinder FromResolve(object subIdentifier)
        {
            _subBinder.FromResolve(subIdentifier);
            return new ScopeNonLazyBinder(_info);
        }

        public ScopeArgNonLazyBinder FromFactory(Type factoryType)
        {
            _subBinder.FromFactory(factoryType);
            return this;
        }

#if !NOT_UNITY3D

        public ScopeArgNonLazyBinder FromNewComponentOn(GameObject gameObject)
        {
            _subBinder.FromNewComponentOn(gameObject);
            return this;
        }

        public ArgNonLazyBinder FromNewComponentSibling()
        {
            _subBinder.FromNewComponentSibling();
            return this;
        }

        public NameTransformScopeArgNonLazyBinder FromNewComponentOnNewGameObject()
        {
            var gameObjectInfo = new GameObjectCreationParameters();
            _subBinder.FromNewComponentOnNewGameObject(gameObjectInfo);
            return new NameTransformScopeArgNonLazyBinder(_info, gameObjectInfo);
        }

        public NameTransformScopeArgNonLazyBinder FromComponentInNewPrefab(UnityEngine.Object prefab)
        {
            var gameObjectInfo = new GameObjectCreationParameters();
            _subBinder.FromComponentInNewPrefab(prefab, gameObjectInfo);
            return new NameTransformScopeArgNonLazyBinder(_info, gameObjectInfo);
        }

        public NameTransformScopeArgNonLazyBinder FromComponentInNewPrefabResource(string resourcePath)
        {
            var gameObjectInfo = new GameObjectCreationParameters();
            _subBinder.FromComponentInNewPrefabResource(resourcePath, gameObjectInfo);
            return new NameTransformScopeArgNonLazyBinder(_info, gameObjectInfo);
        }

        public ScopeNonLazyBinder FromResource(string resourcePath)
        {
            _subBinder.FromResource(resourcePath);
            return new ScopeNonLazyBinder(_info);
        }

#endif

        public ScopeArgNonLazyBinder FromMethodUntyped(Func<InjectContext, object> method)
        {
            _subBinder.FromMethodUntyped(method);
            return this;
        }
    }
}
