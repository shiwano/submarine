using UnityEngine;
using Zenject;

namespace Submarine.Battle
{
    [RequireComponent(typeof(SubmarineView))]
    public class SubmarineInstaller : ActorInstaller<SubmarineFacade, SubmarineView>
    {
        [Inject]
        bool isPlayerSubmarine;

        public override void InstallBindings()
        {
            base.InstallBindings();

            if (isPlayerSubmarine)
            {
                Container.BindInterfacesAndSelfTo<PlayerSubmarineMediator>().AsSingle();
            }
        }
    }
}
