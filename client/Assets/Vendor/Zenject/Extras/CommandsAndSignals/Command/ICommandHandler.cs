using System;

namespace Zenject.Commands
{
    public interface ICommandHandlerBase
    {
    }

    // Zero params
    public interface ICommandHandler : ICommandHandlerBase
    {
        void Execute();
    }

    // One param
    public interface ICommandHandler<TParam1> : ICommandHandlerBase
    {
        void Execute(TParam1 p1);
    }

    // Two params
    public interface ICommandHandler<TParam1, TParam2> : ICommandHandlerBase
    {
        void Execute(TParam1 p1, TParam2 p2);
    }

    // Three params
    public interface ICommandHandler<TParam1, TParam2, TParam3> : ICommandHandlerBase
    {
        void Execute(TParam1 p1, TParam2 p2, TParam3 p3);
    }

    // Four params
    public interface ICommandHandler<TParam1, TParam2, TParam3, TParam4> : ICommandHandlerBase
    {
        void Execute(TParam1 p1, TParam2 p2, TParam3 p3, TParam4 p4);
    }

    // Five params
    public interface ICommandHandler<TParam1, TParam2, TParam3, TParam4, TParam5> : ICommandHandlerBase
    {
        void Execute(TParam1 p1, TParam2 p2, TParam3 p3, TParam4 p4, TParam5 p5);
    }

    // Six params
    public interface ICommandHandler<TParam1, TParam2, TParam3, TParam4, TParam5, TParam6> : ICommandHandlerBase
    {
        void Execute(TParam1 p1, TParam2 p2, TParam3 p3, TParam4 p4, TParam5 p5, TParam6 p6);
    }
}
