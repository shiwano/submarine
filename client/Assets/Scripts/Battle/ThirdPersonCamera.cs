using UnityEngine;
using Zenject;

namespace Submarine.Battle
{
    public class ThirdPersonCamera : ITickable
    {
        readonly Transform camera;
        readonly Quaternion cameraStartRotation;

        Transform target;

        Vector3 CameraDistance
        {
            get { return -target.forward * 18.5f + target.right * 1.8f + target.up * 8f; }
        }

        public ThirdPersonCamera()
        {
            this.camera = Camera.main.transform;
            cameraStartRotation = camera.transform.rotation;
        }

        public void SetTarget(Transform target)
        {
            this.target = target;
        }

        public void Tick()
        {
            if (target != null)
            {
                camera.position = target.position + CameraDistance;
                camera.rotation = target.rotation * cameraStartRotation;
            }
        }
    }
}
