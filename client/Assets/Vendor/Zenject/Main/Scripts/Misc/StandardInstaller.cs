using System;
using System.Linq;

namespace Zenject
{
    public class StandardInstaller<TRoot> : Installer
        where TRoot : IFacade
    {
        // Install basic functionality for most unity apps
        public override void InstallBindings()
        {
            Container.Bind<IFacade>().ToSingle<TRoot>();
            Container.Bind<TRoot>().ToSingle();

            Container.Bind<TickableManager>().ToSingle();
            Container.Bind<InitializableManager>().ToSingle();
            Container.Bind<DisposableManager>().ToSingle();
        }
    }

    public class StandardInstaller : StandardInstaller<Facade>
    {
    }
}
