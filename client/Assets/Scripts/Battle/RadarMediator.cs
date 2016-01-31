using Zenject;

namespace Submarine.Battle
{
    public class RadarMediator : IInitializable
    {
        [Inject]
        RadarView view;

        public void Initialize()
        {
        }
    }
}
