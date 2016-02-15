using Zenject;

namespace Submarine.Battle
{
    public class SubmarineFacade : Facade
    {
        public class Factory : FacadeFactory<SubmarineFacade> { }

        [Inject]
        SubmarineView view;
    }
}
