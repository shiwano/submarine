using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ThirdPersonCamera : ITickable
    {
        private readonly Transform camera;
        private readonly Quaternion cameraStartRotation;

        private Transform target;

        public Vector3 CameraDistance
        {
            get { return -target.forward * 185f + target.right * 18f + target.up * 80f; }
        }

        public ThirdPersonCamera([Inject("MainCamera")] Camera camera)
        {
            this.camera = camera.transform;
            cameraStartRotation = camera.transform.rotation;
        }

        public void SetTarget(Transform target)
        {
            this.target = target;
        }

        public void Tick()
        {
            if (target == null)
            {
                return;
            }

            camera.position = target.position + CameraDistance;
            camera.rotation = target.rotation * cameraStartRotation;
        }
    }
}
