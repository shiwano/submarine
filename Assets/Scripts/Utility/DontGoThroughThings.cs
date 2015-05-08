using UnityEngine;

namespace Submarine
{
    /// <summary>
    /// Imported from http://wiki.unity3d.com/index.php?title=DontGoThroughThings
    /// </summary>
    public class DontGoThroughThings : MonoBehaviour
    {
        public LayerMask layerMask;    // make sure we aren't in this layer
        public float skinWidth = 0.1f; // probably doesn't need to be changed

        float minimumExtent;
        float partialExtent;
        float sqrMinimumExtent;
        Vector3 previousPosition;
        Rigidbody myRigidbody;
        Collider myCollider;

        void Awake()
        { 
            myRigidbody = GetComponent<Rigidbody>();
            myCollider = GetComponent<Collider>();
            previousPosition = myRigidbody.position;
            minimumExtent = Mathf.Min(Mathf.Min(myCollider.bounds.extents.x, myCollider.bounds.extents.y), myCollider.bounds.extents.z);
            partialExtent = minimumExtent * (1.0f - skinWidth);
            sqrMinimumExtent = minimumExtent * minimumExtent;
        } 
     
        void FixedUpdate() 
        { 
            // have we moved more than our minimum extent?
            Vector3 movementThisStep = myRigidbody.position - previousPosition;
            float movementSqrMagnitude = movementThisStep.sqrMagnitude;

            if (movementSqrMagnitude > sqrMinimumExtent)
            {
                float movementMagnitude = Mathf.Sqrt(movementSqrMagnitude);
                RaycastHit hitInfo;

                // check for obstructions we might have missed
                if (Physics.Raycast(previousPosition, movementThisStep, out hitInfo, movementMagnitude, layerMask.value))
                {
                    myRigidbody.position = hitInfo.point - (movementThisStep/movementMagnitude) * partialExtent;
                }
            }

            previousPosition = myRigidbody.position;
        }
    }
}
