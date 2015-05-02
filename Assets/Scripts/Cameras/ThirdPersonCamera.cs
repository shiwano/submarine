using UnityEngine;
using Zenject;

namespace Submarine
{
    public class ThirdPersonCamera : ITickable, IInitializable
    {
        private readonly Transform camera;
        private Vector3 offset;
        private Transform target;

        public ThirdPersonCamera(
            [Inject("MainCamera")] Camera camera)
        {
            this.camera = camera.transform;
        }

        public void Initialize()
        {
            UpdateCameraPosition(Vector3.zero);
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

            UpdateCameraPosition(target.position);
        }

        void UpdateCameraPosition(Vector3 position)
        {
            var dest = position + offset;
            camera.position = new Vector3(dest.x, camera.position.y, dest.z);
        }
    }
}
