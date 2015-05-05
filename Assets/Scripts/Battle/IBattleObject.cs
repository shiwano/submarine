using System;
using Zenject;

namespace Submarine
{
    public enum BattleObjectType
    {
        Submarine = 0,
        Torpedo,
    }

    public interface IBattleObject : IInitializable, ITickable, IDisposable
    {
        BattleObjectType Type { get; }
        Photon.MonoBehaviour PhotonMonoBehaviour { get; }
    }
}
