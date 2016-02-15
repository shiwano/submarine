using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class SubmarineFacade : Facade, IActor
    {
        public class Factory : FacadeFactory<Type.Battle.Actor, SubmarineFacade> { }

        [Inject]
        SubmarineView view;
        [Inject]
        Type.Battle.Actor actor;

        public SubmarineView View { get { return view; } }
    }
}
