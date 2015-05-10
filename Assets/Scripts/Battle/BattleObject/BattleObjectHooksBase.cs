using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public interface IBattleObjectHooks : IDisposable
    {
        BattleObjectType Type { get; }
        bool IsMine { get; }
        int ViewId { get; }
        Transform transform { get; }
    }

    [RequireComponent(
        typeof(PhotonView),
        typeof(BoxCollider),
        typeof(Rigidbody)
    )]
    public abstract class BattleObjectHooksBase : Photon.MonoBehaviour, IBattleObjectHooks
    {
        public static event Action<IBattleObjectHooks> CreatedViaNetwork = delegate {};
        public static event Action<IBattleObjectHooks> DestroyedViaNetwork = delegate {};

        BattleService battleService;

        public abstract BattleObjectType Type { get; }
        public bool IsMine { get { return photonView.isMine; } }
        public int ViewId { get { return photonView.viewID; } }

        public Rigidbody Rigidbody { get; private set; }

        protected abstract void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info);

        [PostInject]
        public void Initialize(BattleService battleService)
        {
            this.battleService = battleService;
        }

        public virtual void Dispose()
        {
            if (IsMine && gameObject != null)
            {
                battleService.DestroyPhotonView(gameObject);
            }
        }

        protected virtual void Awake()
        {
            Rigidbody = GetComponent<Rigidbody>();

            if (!IsMine)
            {
                BattleObjectHooksBase.CreatedViaNetwork(this);
            }
        }

        protected virtual void OnDestroy()
        {
            if (!IsMine)
            {
                BattleObjectHooksBase.DestroyedViaNetwork(this);
            }
        }
    }
}
