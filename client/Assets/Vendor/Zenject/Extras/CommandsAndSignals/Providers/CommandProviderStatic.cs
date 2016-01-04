using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using System.Linq;

namespace Zenject.Commands
{
    public class CommandProviderStatic<TCommand, TAction> : ProviderBase
        where TCommand : ICommand
    {
        readonly DiContainer _container;
        readonly Func<TAction> _methodGetter;

        public CommandProviderStatic(
            DiContainer container, Func<TAction> methodGetter)
        {
            _methodGetter = methodGetter;
            _container = container;
        }

        public override Type GetInstanceType()
        {
            return typeof(TCommand);
        }

        public override object GetInstance(InjectContext context)
        {
            var obj = _container.Instantiate<TCommand>(_methodGetter());
            Assert.That(obj != null);
            return obj;
        }

        public override IEnumerable<ZenjectResolveException> ValidateBinding(InjectContext context)
        {
            return _container.ValidateObjectGraph<TCommand>(context, typeof(TAction));
        }
    }
}

