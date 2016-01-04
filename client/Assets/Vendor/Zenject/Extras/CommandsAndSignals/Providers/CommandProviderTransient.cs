using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using System.Linq;

namespace Zenject.Commands
{
    public abstract class CommandProviderTransient<TCommand, THandler, TAction>
        : CommandProviderBase<TCommand, TAction>
        where TCommand : ICommand
    {
        public CommandProviderTransient(DiContainer container)
            : base(container)
        {
        }

        protected THandler CreateHandler(InjectContext c)
        {
            var newContext = new InjectContext(
                c.Container, typeof(THandler), null, false, c.ObjectType,
                c.ObjectInstance, c.MemberName, c.ParentContext, c.ConcreteIdentifier,
                null, c.LocalOnly);

            return Container.InstantiateExplicit<THandler>(new List<TypeValuePair>(), newContext);
        }

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            return base.ValidateBinding(context)
                .Concat(Container.ValidateObjectGraph<THandler>(context));
        }
    }
}


