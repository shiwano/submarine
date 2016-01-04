using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using ModestTree.Util;
using System.Linq;

namespace Zenject.Commands
{
    // Zero params
    public class CommandProviderMethodSingle<TCommand, THandler>
        : CommandProviderSingle<TCommand, THandler, Action>
        where TCommand : Command
    {
        readonly Func<THandler, Action> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container, Func<THandler, Action> methodGetter, ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override Action GetCommandAction(InjectContext context)
        {
            return () =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)();
            };
        }
    }

    // One param
    public class CommandProviderMethodSingle<TCommand, THandler, TParam1>
        : CommandProviderSingle<TCommand, THandler, Action<TParam1>>
        where TCommand : Command<TParam1>
    {
        readonly Func<THandler, Action<TParam1>> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container,
            Func<THandler, Action<TParam1>> methodGetter, ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override Action<TParam1> GetCommandAction(InjectContext context)
        {
            return (p1) =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)(p1);
            };
        }
    }

    // Two params
    public class CommandProviderMethodSingle<TCommand, THandler, TParam1, TParam2>
        : CommandProviderSingle<TCommand, THandler, Action<TParam1, TParam2>>
        where TCommand : Command<TParam1, TParam2>
    {
        readonly Func<THandler, Action<TParam1, TParam2>> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container,
            Func<THandler, Action<TParam1, TParam2>> methodGetter,
            ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override Action<TParam1, TParam2> GetCommandAction(InjectContext context)
        {
            return (p1, p2) =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)(p1, p2);
            };
        }
    }

    // Three params
    public class CommandProviderMethodSingle<TCommand, THandler, TParam1, TParam2, TParam3>
        : CommandProviderSingle<TCommand, THandler, Action<TParam1, TParam2, TParam3>>
        where TCommand : Command<TParam1, TParam2, TParam3>
    {
        readonly Func<THandler, Action<TParam1, TParam2, TParam3>> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container,
            Func<THandler, Action<TParam1, TParam2, TParam3>> methodGetter,
            ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override Action<TParam1, TParam2, TParam3> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3) =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)(p1, p2, p3);
            };
        }
    }

    // Four params
    public class CommandProviderMethodSingle<TCommand, THandler, TParam1, TParam2, TParam3, TParam4>
        : CommandProviderSingle<TCommand, THandler, Action<TParam1, TParam2, TParam3, TParam4>>
        where TCommand : Command<TParam1, TParam2, TParam3, TParam4>
    {
        readonly Func<THandler, Action<TParam1, TParam2, TParam3, TParam4>> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container,
            Func<THandler, Action<TParam1, TParam2, TParam3, TParam4>> methodGetter,
            ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override Action<TParam1, TParam2, TParam3, TParam4> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3, p4) =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)(p1, p2, p3, p4);
            };
        }
    }

    // Five params
    public class CommandProviderMethodSingle<TCommand, THandler, TParam1, TParam2, TParam3, TParam4, TParam5>
        : CommandProviderSingle<TCommand, THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5>>
        where TCommand : Command<TParam1, TParam2, TParam3, TParam4, TParam5>
    {
        readonly Func<THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5>> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container,
            Func<THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5>> methodGetter,
            ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3, p4, p5) =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)(p1, p2, p3, p4, p5);
            };
        }
    }

    // Six params
    public class CommandProviderMethodSingle<TCommand, THandler, TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>
        : CommandProviderSingle<TCommand, THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>>
        where TCommand : Command<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>
    {
        readonly Func<THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>> _methodGetter;

        public CommandProviderMethodSingle(
            DiContainer container,
            Func<THandler, ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6>> methodGetter,
            ProviderBase singletonProvider)
            : base(container, singletonProvider)
        {
            _methodGetter = methodGetter;
        }

        protected override ModestTree.Util.Action<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6> GetCommandAction(InjectContext context)
        {
            return (p1, p2, p3, p4, p5, p6) =>
            {
                var singleton = GetSingleton(context);
                Assert.IsNotNull(singleton);
                _methodGetter(singleton)(p1, p2, p3, p4, p5, p6);
            };
        }
    }
}
