using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class SubmarineFacade : ActorFacade
    {
        public class Factory : FacadeFactory<Type.Battle.Actor, bool, SubmarineFacade> { }

        [Inject]
        SubmarineView view;

        public override bool WillIgnoreMotorDirection
        {
            get { return IsMine; }
        }
    }
}
