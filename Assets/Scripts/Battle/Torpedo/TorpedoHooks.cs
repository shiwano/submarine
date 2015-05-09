using UnityEngine;
using System;

namespace Submarine
{
    public class TorpedoHooks : BattleObjectHooksBase
    {
        Vector3 receivedPosition = Vector3.zero;
        Quaternion receivedRotation = Quaternion.identity;

        public override BattleObjectType Type { get { return BattleObjectType.Torpedo; } }

        public event Action<int?> Striked = delegate {};

        public void Accelerate(Vector3 force)
        {
            Rigidbody.AddForce(force, ForceMode.Force);
        }

        void OnCollisionEnter(Collision collision)
        {
            if (IsMine)
            {
                var submarineHooks = collision.gameObject.GetComponent<SubmarineHooks>();
                if (submarineHooks != null && !submarineHooks.photonView.isMine)
                {
                    Striked(submarineHooks.photonView.viewID);
                }
                else
                {
                    Striked(null);
                }
            }
        }

        void Update()
        {
            if (!IsMine)
            {
                transform.position = Vector3.Lerp(transform.position, receivedPosition, Time.deltaTime * 5);
                transform.rotation = Quaternion.Lerp(transform.rotation, receivedRotation, Time.deltaTime * 5);
            }
        }

        protected override void OnPhotonSerializeView(PhotonStream stream, PhotonMessageInfo info)
        {
            if (stream.isWriting)
            {
                stream.SendNext(transform.position);
                stream.SendNext(transform.rotation);
            }
            else
            {
                receivedPosition = (Vector3)stream.ReceiveNext();
                receivedRotation = (Quaternion)stream.ReceiveNext();
            }
        }
    }
}
