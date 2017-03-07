using System;

namespace Submarine
{
    public interface IClock
    {
        DateTime Now { get; }
    }
}
