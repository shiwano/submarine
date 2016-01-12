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
        public static event Action<IBattleObjectHooks> Created = delegate {};
        public static event Action<IBattleObjectHooks> Destroyed = delegate {};

        OldBattleService battleService;

        public abstract BattleObjectType Type { get; }
        public bool IsMine { get { return photonView.isMine; } }
        public int ViewId { get { return photonView.viewID; } }

        public Rigidbody Rigidbody { get; private set; }

        protected abstract void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info);

        [PostInject]
        public void Initialize(OldBattleService battleService)
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
            BattleObjectHooksBase.Created(this);
        }

        protected virtual void OnDestroy()
        {
            BattleObjectHooksBase.Destroyed(this);
        }
    }
}
