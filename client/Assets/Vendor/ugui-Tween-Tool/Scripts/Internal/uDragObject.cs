using UnityEngine;
using UnityEngine.EventSystems;
using System.Collections;

namespace uTools {
	[AddComponentMenu("uTools/Internal/Drag Object(uTools)")]
	public class uDragObject : MonoBehaviour, IDragHandler {

		public RectTransform target;

		RectTransform cacheTarget {
			get {
				if (target == null) {
					target = GetComponent<RectTransform>();
				}
				return target;
			}
		}

		// Use this for initialization
		void Start () {		

		}
		
		// Update is called once per frame
		void Update () {
			
		}

		public void OnDrag (PointerEventData eventData) {
			Vector3 from = cacheTarget.localPosition;
			Vector3 to = from + new Vector3 (eventData.delta.x, eventData.delta.y, 0);
			uTweenPosition.Begin (gameObject, from, to, .02f);//.easeType = EaseType.easeInBack;
			//cacheTarget.localPosition += new Vector3 (eventData.delta.x, eventData.delta.y, 0);
		}
	}
}