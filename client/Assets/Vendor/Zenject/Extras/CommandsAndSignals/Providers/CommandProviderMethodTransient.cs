using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using ModestTree.Util;
using System.Linq;

namespace Zenject.Commands
{
    // Common
    public abstract class CommandProviderMethodTransientBase<TCommand, THandler, TAction>
        : CommandProviderTransient<TCommand, THandler, TAction>
        where TCommand : ICommand
    {
        readonly Func<THandler, TAction> _methodGetter;

        public CommandProviderMethodTransientBase(
            DiContainer container, Func<THandler, TAction> methodGetter)
            : base(container)
        {
            _methodGetter = methodGetter;
        }

        protected TAction GetHandlerMethod(InjectContext context)
        {
            return _methodGetter(CreateHandler(context));
        }
    }

    // Zero Parameters
    public class CommandProviderMethodTransient<TCommand, THandler>
        : CommandProviderMethodTransientBase<TCommand, THandler, Action>
        where TCommand : Command
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, Action> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override Action GetCommandAction(InjectContext context)
        {
            return () =>
            {
                GetHandlerMethod(context)();
            };
        }
    }

    // One Parameter
    public class CommandProviderMethodTransient<TCommand, THandler, TParam1>
        : CommandProviderMethodTransientBase<TCommand, THandler, Action<TParam1>>
        where TCommand : Command<TParam1>
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, Action<TParam1>> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override Action<TParam1> GetCommandAction(InjectContext context)
        {
            return (p1) =>
            {
                GetHandlerMethod(context)(p1);
            };
        }
    }

    // Two Parameters
    public class CommandProviderMethodTransient<TCommand, THandler, TParam1, TParam2>
        : CommandProviderMethodTransientBase<TCommand, THandler, Action<TParam1, TParam2>>
        where TCommand : Command<TParam1, TParam2>
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, Action<TParam1, TParam2>> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override Action<TParam1, TParam2> GetCommandAction(InjectContext context)
        {
            return (p1, p2) =>
            {
                GetHandlerMethod(context)(p1, p2);
            };
        }
    }

    // Three Parameters
    public class CommandProviderMethodTransient<TCommand, THandler, TParam1, TParam2, TParam3>
        : CommandProviderMethodTransientBase<TCommand, THandler, Action<TParam1, TParam2, TParam3>>
        where TCommand : Command<TParam1, TParam2, TParam3>
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, Action<TParam1, TParam2, TParam3>> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override Action<TParam1, TParam2, TParam3> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3) =>
            {
                GetHandlerMethod(context)(p1, p2, p3);
            };
        }
    }

    // Four Parameters
    public class CommandProviderMethodTransient<TCommand, THandler, TParam1, TParam2, TParam3, TParam4>
        : CommandProviderMethodTransientBase<TCommand, THandler, Action<TParam1, TParam2, TParam3, TParam4>>
        where TCommand : Command<TParam1, TParam2, TParam3, TParam4>
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, Action<TParam1, TParam2, TParam3, TParam4>> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override Action<TParam1, TParam2, TParam3, TParam4> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3, p4) =>
            {
                GetHandlerMethod(context)(p1, p2, p3, p4);
            };
        }
    }

    // Five Parameters
    public class CommandProviderMethodTransient<TCommand, THandler, TParam1, TParam2, TParam3, TParam4, TParam5>
        : CommandProviderMethodTransientBase<TCommand, THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5>>
        where TCommand : Command<TParam1, TParam2, TParam3, TParam4, TParam5>
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5>> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3, p4, p5) =>
            {
                GetHandlerMethod(context)(p1, p2, p3, p4, p5);
            };
        }
    }

    // Six Parameters
    public class CommandProviderMethodTransient<TCommand, THandler, TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>
        : CommandProviderMethodTransientBase<TCommand, THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>>
        where TCommand : Command<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>
    {
        public CommandProviderMethodTransient(
            DiContainer container, Func<THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>> methodGetter)
            : base(container, methodGetter)
        {
        }

        protected override ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3, p4, p5, p6) =>
            {
                GetHandlerMethod(context)(p1, p2, p3, p4, p5, p6);
            };
        }
    }
}

