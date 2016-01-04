using System;
using System.Collections.Generic;

namespace Zenject
{
    public interface IFacade : IInitializable, IDisposable, ITickable, ILateTickable, IFixedTickable
    {
    }
}
