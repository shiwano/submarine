using System;
using System.Collections.Generic;

namespace Zenject
{
    public interface IFacade : IDisposable
    {
        void Initialize();
        void Tick();
        void LateTick();
        void FixedTick();
    }
}
