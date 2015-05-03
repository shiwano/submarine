using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ThirdPersonCamera : ITickable
    {
        private readonly Transform camera;
        private readonly Vector3 cameraStartPosition;
        private readonly Quaternion cameraStartRotation;

        private Transform target;

        public ThirdPersonCamera([Inject("MainCamera")] Camera camera)
        {
            this.camera = camera.transform;
            cameraStartPosition = camera.transform.position;
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

            var dest = -target.forward * cameraStartPosition.magnitude;
            camera.position = new Vector3(
                target.position.x + dest.x,
                cameraStartPosition.y,
                target.position.z + dest.z
            );
            camera.rotation = target.rotation * cameraStartRotation;
        }
    }
}
