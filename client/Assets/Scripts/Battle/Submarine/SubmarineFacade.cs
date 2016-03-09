using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class SubmarineFacade : ActorFacade
    {
        public class Factory : FacadeFactory<Type.Battle.Actor, SubmarineFacade> { }

        [Inject]
        SubmarineView view;

        public void Turn(float rate)
        {
            view.Turn(rate);
        }
    }
}
