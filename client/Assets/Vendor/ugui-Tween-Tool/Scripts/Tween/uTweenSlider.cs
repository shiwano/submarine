using UnityEngine;
using UnityEngine.UI;
using System.Collections;

namespace uTools {
	[AddComponentMenu("uTools/Tween/Tween Slider(uTools)")]	
	public class uTweenSlider : uTweenValue {

		private Slider mSlider;
		public Slider cacheSlider {
			get {
				mSlider = GetComponent<Slider>();
				if (mSlider == null) {
					Debug.LogError("'uTweenSlider' can't find 'Slider'");
				}
				return mSlider;
			}
		}

		/// <summary>
		/// The need carry.
		/// when is true, value==1 equal value=0
		/// </summary>
		public bool NeedCarry = false;

		public float sliderValue {
			set {
				if (NeedCarry) {
					cacheSlider.value = (value>=1)?value - Mathf.Floor(value) : value;
				}
				else {
					cacheSlider.value = (value>1)?value - Mathf.Floor(value) : value;
				}
			}
		}

		protected override void ValueUpdate (float value, bool isFinished)
		{
			this.sliderValue = value;
		}

		public static uTweenSlider Begin(Slider slider, float duration, float delay, float from, float to) {
			uTweenSlider comp = uTweener.Begin<uTweenSlider>(slider.gameObject, duration);
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
