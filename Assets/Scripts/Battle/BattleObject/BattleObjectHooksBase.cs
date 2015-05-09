using UnityEngine;
using System;
using Zenject;

namespace Submarine
{
    public interface IBattleObjectHooks : IDisposable
    {
        BattleObjectType Type { get; }
        bool IsMine { get; }
    }

    public abstract class BattleObjectHooksBase : Photon.MonoBehaviour, IBattleObjectHooks
    {
        BattleService battleService;

        public abstract BattleObjectType Type { get; }

        public bool IsMine { get { return photonView.isMine; } }
        public Rigidbody Rigidbody { get; private set; }

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
            BattleEvent.BattleObjectHooksCreated(this);
        }

        protected virtual void OnDestroy()
        {
            BattleEvent.BattleObjectHooksDestroyed(this);
        }
    }
}
