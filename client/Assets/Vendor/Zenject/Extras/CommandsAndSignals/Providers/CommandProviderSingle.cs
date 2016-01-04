using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using System.Linq;

namespace Zenject.Commands
{
    public abstract class CommandProviderSingle<TCommand, THandler, TAction>
        : CommandProviderBase<TCommand, TAction>
        where TCommand : ICommand
    {
        readonly ProviderBase _singletonProvider;

        public CommandProviderSingle(
            DiContainer container, ProviderBase singletonProvider)
            : base(container)
        {
            _singletonProvider = singletonProvider;
        }

        public override void Dispose()
        {
            _singletonProvider.Dispose();
        }

        protected THandler GetSingleton(InjectContext c)
        {
            var newContext = new InjectContext(
                c.Container, typeof(THandler), null, false, c.ObjectType,
                c.ObjectInstance, c.MemberName, c.ParentContext, c.ConcreteIdentifier,
                null, c.LocalOnly);

            return (THandler)_singletonProvider.GetInstance(newContext);
        }

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            return base.ValidateBinding(context)
                .Concat(_singletonProvider.ValidateBinding(context));
        }
    }
}

