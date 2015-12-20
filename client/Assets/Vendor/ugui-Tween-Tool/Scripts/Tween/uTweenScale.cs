using UnityEngine;
using UnityEngine.UI;
using System.Collections;

namespace uTools {
	[AddComponentMenu("uTools/Tween/Tween Scale(uTools)")]
	
	public class uTweenScale : uTweener {

		public Vector3 from = Vector3.zero;
		public Vector3 to = Vector3.one;

		RectTransform mRectTransform;

		public RectTransform cachedRectTransform { get { if (mRectTransform == null) mRectTransform = GetComponent<RectTransform>(); return mRectTransform;}}
		public Vector3 value {
			get { return cachedRectTransform.localScale;}
			set { cachedRectTransform.localScale = value;}
		}

		protected override void OnUpdate (float factor, bool isFinished)
		{
			value = from + factor * (to - from);
		}

		public static uTweenScale Begin(GameObject go, Vector3 from, Vector3 to, float duration = 1f, float delay = 0f) {
			uTweenScale comp = uTweener.Begin<uTweenScale>(go, duration);
			comp.from = from;
			comp.to = to;
			comp.duration = duration;
			comp.delay = delay;
			if (duration <= 0) {
				comp.Sample(1, true);
				comp.enabled = false;
			}
			return comp;
		}
	}
}
