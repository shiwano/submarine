using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ThirdPersonCamera : ITickable, IInitializable
    {
        private readonly Transform camera;
        private readonly Quaternion cameraStartRotation;

        private Vector3 offset;
        private Transform target;

        public ThirdPersonCamera([Inject("MainCamera")] Camera camera)
        {
            this.camera = camera.transform;
            cameraStartRotation = camera.transform.rotation;
        }

        public void Initialize()
        {
            UpdateCameraTransform(Vector3.zero, Quaternion.identity);
            RaycastHit hit;

            if (Physics.Raycast(camera.position, Vector3.forward, out hit))
            {
                offset = hit.transform.position;
            }
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

            UpdateCameraTransform(target.position, target.rotation);
        }

        void UpdateCameraTransform(Vector3 position, Quaternion rotation)
        {
            var dest = position + offset;
            camera.position = new Vector3(dest.x, camera.position.y, dest.z);
            camera.rotation = rotation * cameraStartRotation;
        }
    }
}
