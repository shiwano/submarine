using UnityEngine;
using System;

namespace Submarine
{
    public interface IBattleObjectHooks : IDisposable
    {
        BattleObjectType Type { get; }
        bool IsMine { get; }
    }

    public abstract class BattleObjectHooksBase : Photon.MonoBehaviour, IBattleObjectHooks
    {
        public abstract BattleObjectType Type { get; }

        public bool IsMine { get { return photonView.isMine; } }
        public Rigidbody Rigidbody { get; private set; }

        public virtual void Dispose()
        {
            if (IsMine && gameObject != null)
            {
                PhotonNetwork.Destroy(gameObject);
            }
        }

        protected virtual void Awake()
        {
            Rigidbody = GetComponent<Rigidbody>();
            BattleEvent.BattleObjectHooksCreated(this);
        }

        protected virtual void OnDestroy()
        {
            BattleEvent.BattleObjectHooksDestroyed(this);
        }
    }
}
