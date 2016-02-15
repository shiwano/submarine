using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class SubmarineFacade : ActorFacade
    {
        public class Factory : FacadeFactory<Type.Battle.Actor, SubmarineFacade> { }

        [Inject]
        SubmarineView view;

        public SubmarineView View { get { return view; } }
    }
}
