using Zenject;

namespace Submarine
{
    public abstract class MediatorBase<TView> where TView : IView
    {
        [Inject]
        protected TView view;
    }
}
