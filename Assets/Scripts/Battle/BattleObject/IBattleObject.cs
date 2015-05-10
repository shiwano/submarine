using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public enum BattleObjectType
    {
        Submarine = 0,
        Torpedo,
        Decoy,
        Lookout,
    }

    public interface IBattleObject : IInitializable, ITickable, IDisposable
    {
        BattleObjectType Type { get; }
        IBattleObjectHooks BattleObjectHooks { get; }
        Vector3 Position { get; }
        Vector3 EulerAngles { get; }
        float SearchRange { get; }
        bool IsVisibleFromPlayer { get; }
    }
}
