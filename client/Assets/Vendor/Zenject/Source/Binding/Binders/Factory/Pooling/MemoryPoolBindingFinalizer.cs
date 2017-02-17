using System;
using ModestTree;

namespace Zenject
{
    public class MemoryPoolBindingFinalizer<TContract> : ProviderBindingFinalizer
    {
        readonly MemoryPoolBindInfo _poolBindInfo;
        readonly FactoryBindInfo _factoryBindInfo;

        public MemoryPoolBindingFinalizer(
            BindInfo bindInfo, FactoryBindInfo factoryBindInfo, MemoryPoolBindInfo poolBindInfo)
            : base(bindInfo)
        {
            // Note that it doesn't derive from MemoryPool<TContract>
            // when used with To<>, so we can only check IMemoryPoolBase
            Assert.That(factoryBindInfo.FactoryType.DerivesFrom<IMemoryPool>());

            _factoryBindInfo = factoryBindInfo;
            _poolBindInfo = poolBindInfo;
        }

        protected override void OnFinalizeBinding(DiContainer container)
        {
            var provider = _factoryBindInfo.ProviderFunc(container);

            RegisterProviderForAllContracts(
                container,
                new CachedProvider(
                    new TransientProvider(
                        _factoryBindInfo.FactoryType,
                        container,
                        InjectUtil.CreateArgListExplicit(
                            typeof(TContract), provider, _poolBindInfo.InitialSize, _poolBindInfo.ExpandMethod),
                        null,
                        BindInfo.ContextInfo)));
        }
    }
}

