using Zenject;

namespace Submarine
{
    public static class ZenjectExtensions
    {
        public static void BindMediatorAndViewAsSingle<TMediator, TView>(this DiContainer container, TView view)
            where TMediator : MediatorBase<TView>
            where TView : IView
        {
            container.BindInterfacesAndSelfTo<TMediator>().AsSingle();
            container.Bind<TView>().FromInstance(view).AsSingle();
        }
    }
}
