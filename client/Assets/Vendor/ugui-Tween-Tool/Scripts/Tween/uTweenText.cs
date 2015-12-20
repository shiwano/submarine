using UnityEngine;
using UnityEngine.UI;
using System.Collections;

namespace uTools {
	[AddComponentMenu("uTools/Tween/Tween Text(uTools)")]	
	
	public class uTweenText : uTweenValue {

		private Text mText;
		public Text cacheText {
			get {
				mText = GetComponent<Text>();
				if (mText == null) {
					Debug.LogError("'uTweenText' can't find 'Text'");
				}
				return mText;
			}
		}

		/// <summary>
		/// number after the digit point
		/// </summary>
		public int digits;

		protected override void ValueUpdate (float value, bool isFinished)
		{
			cacheText.text = (System.Math.Round(value, digits)).ToString();
		}

		public static uTweenText Begin(Text label, float duration, float delay, float from, float to) {
			uTweenText comp = uTweener.Begin<uTweenText>(label.gameObject, duration);
			comp.from = from;
			comp.to = to;
			comp.delay = delay;
			
			if (duration <=0) {
				comp.Sample(1, true);
				comp.enabled = false;
			}
			return comp;
		}
	}
}
