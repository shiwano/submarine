using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ThirdPersonCamera : ITickable
    {
        readonly Transform camera;
        readonly Quaternion cameraStartRotation;

        Transform target;

        public Vector3 CameraDistance
        {
            get { return -target.forward * 92.5f + target.right * 9f + target.up * 40f; }
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
            if (target != null)
            {
                camera.position = target.position + CameraDistance;
                camera.rotation = target.rotation * cameraStartRotation;
            }
        }
    }
}
