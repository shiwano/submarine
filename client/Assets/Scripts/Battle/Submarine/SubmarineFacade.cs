using Zenject;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class SubmarineFacade : ActorFacade
    {
        public class Factory : FacadeFactory<Type.Battle.Actor, SubmarineFacade> { }

        [Inject]
        SubmarineView view;

        public double Direction
        {
            get
            {
                var direction = (double)(360f - view.transform.eulerAngles.y + 90f);
                return direction > 360d ? direction - 360d : direction;
            }
        }

        public void Turn(float rate)
        {
            view.Turn(rate);
        }
    }
}
