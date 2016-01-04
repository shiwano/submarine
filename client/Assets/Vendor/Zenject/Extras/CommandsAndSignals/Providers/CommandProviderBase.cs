using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using System.Linq;

namespace Zenject.Commands
{
    public abstract class CommandProviderBase<TCommand, TAction> : ProviderBase
        where TCommand : ICommand
    {
        readonly DiContainer _container;

        public CommandProviderBase(DiContainer container)
        {
            _container = container;
        }

        protected DiContainer Container
        {
            get
            {
                return _container;
            }
        }

        public override Type GetInstanceType()
        {
            return typeof(TCommand);
        }

        public override object GetInstance(InjectContext context)
        {
            var obj = _container.Instantiate<TCommand>(GetCommandAction(context));
            Assert.That(obj != null);
            return obj;
        }

        protected abstract TAction GetCommandAction(InjectContext context);

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            return _container.ValidateObjectGraph<TCommand>(context, typeof(TAction));
        }
    }
}

