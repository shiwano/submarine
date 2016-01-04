using System;
using System.Collections.Generic;
using System.Linq.Expressions;
using ModestTree;
using System.Linq;

namespace Zenject.Commands
{
    public class CommandProviderHandlerTransient<TCommand, THandler>
        : CommandProviderTransient<TCommand, THandler, Action>
        where TCommand : Command
        where THandler : ICommandHandler
    {
        public CommandProviderHandlerTransient(DiContainer container)
            : base(container)
        {
        }

        protected override Action GetCommandAction(InjectContext context)
        {
            return () =>
            {
                CreateHandler(context).Execute();
            };
        }
    }
}


