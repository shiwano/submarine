using UnityEngine;
using Zenject;

namespace Submarine
{
    public abstract class ViewMediatorFactory<TView, TMediator> : IFactory<TView>
        where TView : MonoBehaviour
        where TMediator : new()
    {
        [Inject]
        DiContainer container;

        protected abstract TView CreateView();

        public TView Create()
        {
            var view = CreateView();
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<TView>().ToSingleInstance(view);
            subContainer.Bind<TMediator>().ToSingle();
            return view;
        }
    }

    public abstract class ViewMediatorFactory<TParam1, TView, TMediator> : IFactory<TParam1, TView>
        where TView : MonoBehaviour
        where TMediator : new()
    {
        [Inject]
        DiContainer container;

        protected abstract TView CreateView(TParam1 param1);

        public TView Create(TParam1 param1)
        {
            var view = CreateView(param1);
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<TView>().ToSingleInstance(view);
            subContainer.Bind<TMediator>().ToMethod(ctx => ctx.Container.Instantiate<TMediator>(param1));
            return view;
        }
    }

    public abstract class ViewMediatorFactory<TParam1, TParam2, TView, TMediator> : IFactory<TParam1, TParam2, TView>
        where TView : MonoBehaviour
        where TMediator : new()
    {
        [Inject]
        DiContainer container;

        protected abstract TView CreateView(TParam1 param1, TParam2 param2);

        public TView Create(TParam1 param1, TParam2 param2)
        {
            var view = CreateView(param1, param2);
            var subContainer = container.CreateSubContainer();
            subContainer.Bind<TView>().ToSingleInstance(view);
            subContainer.Bind<TMediator>().ToMethod(ctx => ctx.Container.Instantiate<TMediator>(param1, param2));
            return view;
        }
    }
}
