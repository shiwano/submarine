using UnityEngine;
using System.Collections;

/// <summary>
/// Ease type.
/// </summary>

namespace uTools {
	public enum EaseType{
		none,
		easeInQuad,
		easeOutQuad,
		easeInOutQuad,
		easeInCubic,
		easeOutCubic,
		easeInOutCubic,
		easeInQuart,
		easeOutQuart,
		easeInOutQuart,
		easeInQuint,
		easeOutQuint,
		easeInOutQuint,
		easeInSine,
		easeOutSine,
		easeInOutSine,
		easeInExpo,
		easeOutExpo,
		easeInOutExpo,
		easeInCirc,
		easeOutCirc,
		easeInOutCirc,
		linear,
		spring,
		/* GFX47 MOD START */
		//bounce,
		easeInBounce,
		easeOutBounce,
		easeInOutBounce,
		/* GFX47 MOD END */
		easeInBack,
		easeOutBack,
		easeInOutBack,
		/* GFX47 MOD START */
		//elastic,
		easeInElastic,
		easeOutElastic,
		easeInOutElastic,
		/* GFX47 MOD END */
		punch
	}


	/// <summary>
	/// Loop style.
	/// </summary>
	public enum LoopStyle {
		Once,
		Loop,
		PingPong
	}

	/// <summary>
	/// Ease manager.
	/// </summary>
	public class EaseManager {

		public delegate float EaseDelegate(float start, float end, float t);

		#region ease function	
		private static float linear(float start, float end, float value){
			return Mathf.Lerp(start, end, value);
		}
		
		private static float clerp(float start, float end, float value){
			float min = 0.0f;
			float max = 360.0f;
			float half = Mathf.Abs((max - min) / 2.0f);
			float retval = 0.0f;
			float diff = 0.0f;
			if ((end - start) < -half){
				diff = ((max - start) + end) * value;
				retval = start + diff;
			}else if ((end - start) > half){
				diff = -((max - end) + start) * value;
				retval = start + diff;
			}else retval = start + (end - start) * value;
			return retval;
		}
		
		private static float spring(float start, float end, float value){
			value = Mathf.Clamp01(value);
			value = (Mathf.Sin(value * Mathf.PI * (0.2f + 2.5f * value * value * value)) * Mathf.Pow(1f - value, 2.2f) + value) * (1f + (1.2f * (1f - value)));
			return start + (end - start) * value;
		}
		
		private static float easeInQuad(float start, float end, float value){
			end -= start;
			return end * value * value + start;
		}
		
		private static float easeOutQuad(float start, float end, float value){
			end -= start;
			return -end * value * (value - 2) + start;
		}
		
		private static float easeInOutQuad(float start, float end, float value){
			value /= .5f;
			end -= start;
			if (value < 1) return end / 2 * value * value + start;
			value--;
			return -end / 2 * (value * (value - 2) - 1) + start;
		}
		
		private static float easeInCubic(float start, float end, float value){
			end -= start;
			return end * value * value * value + start;
		}
		
		private static float easeOutCubic(float start, float end, float value){
			value--;
			end -= start;
			return end * (value * value * value + 1) + start;
		}
		
		private static float easeInOutCubic(float start, float end, float value){
			value /= .5f;
			end -= start;
			if (value < 1) return end / 2 * value * value * value + start;
			value -= 2;
			return end / 2 * (value * value * value + 2) + start;
		}
		
		private static float easeInQuart(float start, float end, float value){
			end -= start;
			return end * value * value * value * value + start;
		}
		
		private static float easeOutQuart(float start, float end, float value){
			value--;
			end -= start;
			return -end * (value * value * value * value - 1) + start;
		}
		
		private static float easeInOutQuart(float start, float end, float value){
			value /= .5f;
			end -= start;
			if (value < 1) return end / 2 * value * value * value * value + start;
			value -= 2;
			return -end / 2 * (value * value * value * value - 2) + start;
		}
		
		private static float easeInQuint(float start, float end, float value){
			end -= start;
			return end * value * value * value * value * value + start;
		}
		
		private static float easeOutQuint(float start, float end, float value){
			value--;
			end -= start;
			return end * (value * value * value * value * value + 1) + start;
		}
		
		private static float easeInOutQuint(float start, float end, float value){
			value /= .5f;
			end -= start;
			if (value < 1) return end / 2 * value * value * value * value * value + start;
			value -= 2;
			return end / 2 * (value * value * value * value * value + 2) + start;
		}
		
		private static float easeInSine(float start, float end, float value){
			end -= start;
			return -end * Mathf.Cos(value / 1 * (Mathf.PI / 2)) + end + start;
		}
		
		private static float easeOutSine(float start, float end, float value){
			end -= start;
			return end * Mathf.Sin(value / 1 * (Mathf.PI / 2)) + start;
		}
		
		private static float easeInOutSine(float start, float end, float value){
			end -= start;
			return -end / 2 * (Mathf.Cos(Mathf.PI * value / 1) - 1) + start;
		}
		
		private static float easeInExpo(float start, float end, float value){
			end -= start;
			return end * Mathf.Pow(2, 10 * (value / 1 - 1)) + start;
		}
		
		private static float easeOutExpo(float start, float end, float value){
			end -= start;
			return end * (-Mathf.Pow(2, -10 * value / 1) + 1) + start;
		}
		
		private static float easeInOutExpo(float start, float end, float value){
			value /= .5f;
			end -= start;
			if (value < 1) return end / 2 * Mathf.Pow(2, 10 * (value - 1)) + start;
			value--;
			return end / 2 * (-Mathf.Pow(2, -10 * value) + 2) + start;
		}
		
		private static float easeInCirc(float start, float end, float value){
			end -= start;
			return -end * (Mathf.Sqrt(1 - value * value) - 1) + start;
		}
		
		private static float easeOutCirc(float start, float end, float value){
			value--;
			end -= start;
			return end * Mathf.Sqrt(1 - value * value) + start;
		}
		
		private static float easeInOutCirc(float start, float end, float value){
			value /= .5f;
			end -= start;
			if (value < 1) return -end / 2 * (Mathf.Sqrt(1 - value * value) - 1) + start;
			value -= 2;
			return end / 2 * (Mathf.Sqrt(1 - value * value) + 1) + start;
		}
		
		/* GFX47 MOD START */
		private static float easeInBounce(float start, float end, float value){
			end -= start;
			float d = 1f;
			return end - easeOutBounce(0, end, d-value) + start;
		}
		/* GFX47 MOD END */
		
		/* GFX47 MOD START */
		//private static float bounce(float start, float end, float value){
		private static float easeOutBounce(float start, float end, float value){
			value /= 1f;
			end -= start;
			if (value < (1 / 2.75f)){
				return end * (7.5625f * value * value) + start;
			}else if (value < (2 / 2.75f)){
				value -= (1.5f / 2.75f);
				return end * (7.5625f * (value) * value + .75f) + start;
			}else if (value < (2.5 / 2.75)){
				value -= (2.25f / 2.75f);
				return end * (7.5625f * (value) * value + .9375f) + start;
			}else{
				value -= (2.625f / 2.75f);
				return end * (7.5625f * (value) * value + .984375f) + start;
			}
		}
		/* GFX47 MOD END */
		
		/* GFX47 MOD START */
		private static float easeInOutBounce(float start, float end, float value){
			end -= start;
			float d = 1f;
			if (value < d/2) return easeInBounce(0, end, value*2) * 0.5f + start;
			else return easeOutBounce(0, end, value*2-d) * 0.5f + end*0.5f + start;
		}
		/* GFX47 MOD END */
		
		private static float easeInBack(float start, float end, float value){
			end -= start;
			value /= 1;
			float s = 1.70158f;
			return end * (value) * value * ((s + 1) * value - s) + start;
		}
		
		private static float easeOutBack(float start, float end, float value){
			float s = 1.70158f;
			end -= start;
			value = (value / 1) - 1;
			return end * ((value) * value * ((s + 1) * value + s) + 1) + start;
		}
		
		private static float easeInOutBack(float start, float end, float value){
			float s = 1.70158f;
			end -= start;
			value /= .5f;
			if ((value) < 1){
				s *= (1.525f);
				return end / 2 * (value * value * (((s) + 1) * value - s)) + start;
			}
			value -= 2;
			s *= (1.525f);
			return end / 2 * ((value) * value * (((s) + 1) * value + s) + 2) + start;
		}
		
		private static float punch(float amplitude, float value){
			float s = 9;
			if (value == 0){
				return 0;
			}
			if (value == 1){
				return 0;
			}
			float period = 1 * 0.3f;
			s = period / (2 * Mathf.PI) * Mathf.Asin(0);
			return (amplitude * Mathf.Pow(2, -10 * value) * Mathf.Sin((value * 1 - s) * (2 * Mathf.PI) / period));
		}
		
		/* GFX47 MOD START */
		private static float easeInElastic(float start, float end, float value){
			end -= start;
			
			float d = 1f;
			float p = d * .3f;
			float s = 0;
			float a = 0;
			
			if (value == 0) return start;
			
			if ((value /= d) == 1) return start + end;
			
			if (a == 0f || a < Mathf.Abs(end)){
				a = end;
				s = p / 4;
			}else{
				s = p / (2 * Mathf.PI) * Mathf.Asin(end / a);
			}
			
			return -(a * Mathf.Pow(2, 10 * (value-=1)) * Mathf.Sin((value * d - s) * (2 * Mathf.PI) / p)) + start;
		}		
		/* GFX47 MOD END */
		
		/* GFX47 MOD START */
		//private static float elastic(float start, float end, float value){
		private static float easeOutElastic(float start, float end, float value){
			/* GFX47 MOD END */
			//Thank you to rafael.marteleto for fixing this as a port over from Pedro's UnityTween
			end -= start;
			
			float d = 1f;
			float p = d * .3f;
			float s = 0;
			float a = 0;
			
			if (value == 0) return start;
			
			if ((value /= d) == 1) return start + end;
			
			if (a == 0f || a < Mathf.Abs(end)){
				a = end;
				s = p / 4;
			}else{
				s = p / (2 * Mathf.PI) * Mathf.Asin(end / a);
			}
			
			return (a * Mathf.Pow(2, -10 * value) * Mathf.Sin((value * d - s) * (2 * Mathf.PI) / p) + end + start);
		}		
		
		/* GFX47 MOD START */
		private static float easeInOutElastic(float start, float end, float value){
			end -= start;
			
			float d = 1f;
			float p = d * .3f;
			float s = 0;
			float a = 0;
			
			if (value == 0) return start;
			
			if ((value /= d/2) == 2) return start + end;
			
			if (a == 0f || a < Mathf.Abs(end)){
				a = end;
				s = p / 4;
			}else{
				s = p / (2 * Mathf.PI) * Mathf.Asin(end / a);
			}
			
			if (value < 1) return -0.5f * (a * Mathf.Pow(2, 10 * (value-=1)) * Mathf.Sin((value * d - s) * (2 * Mathf.PI) / p)) + start;
			return a * Mathf.Pow(2, -10 * (value-=1)) * Mathf.Sin((value * d - s) * (2 * Mathf.PI) / p) * 0.5f + end + start;
		}		
		/* GFX47 MOD END */


		public static float EasingFromType(float start, float end, float t, EaseType type){
			switch (type){
			case EaseType.easeInQuad:
				return easeInQuad(start, end, t);
				
			case EaseType.easeOutQuad:
				return easeOutQuad(start, end, t);
				
			case EaseType.easeInOutQuad:
				return easeInOutQuad(start, end, t);
				
			case EaseType.easeInCubic:
				return easeInCubic(start, end, t);
				
			case EaseType.easeOutCubic:
				return easeOutCubic(start, end, t);
				
			case EaseType.easeInOutCubic:
				return easeInOutCubic(start, end, t);
				
			case EaseType.easeInQuart:
				return easeInQuart(start, end, t);
				
			case EaseType.easeOutQuart:
				return easeOutQuart(start, end, t);
				
			case EaseType.easeInOutQuart:
				return easeInOutQuart(start, end, t);
				
			case EaseType.easeInQuint:
				return easeInQuint(start, end, t);
				
			case EaseType.easeOutQuint:
				return easeOutQuint(start, end, t);
				
			case EaseType.easeInOutQuint:
				return easeInOutQuint(start, end, t);
				
			case EaseType.easeInSine:
				return easeInSine(start, end, t);
				
			case EaseType.easeOutSine:
				return easeOutSine(start, end, t);
				
			case EaseType.easeInOutSine:
				return easeInOutSine(start, end, t);
				
			case EaseType.easeInExpo:
				return easeInExpo(start, end, t);
				
			case EaseType.easeOutExpo:
				return easeOutExpo(start, end, t);
				
			case EaseType.easeInOutExpo:
				return easeInOutExpo(start, end, t);
				
			case EaseType.easeInCirc:
				return easeInCirc(start, end, t);
				
			case EaseType.easeOutCirc:
				return easeOutCirc(start, end, t);
				
			case EaseType.easeInOutCirc:
				return easeInOutCirc(start, end, t);
				
			case EaseType.linear:
				return linear(start, end, t);
				
			case EaseType.spring:
				return spring(start, end, t);
				
				/* GFX47 MOD START */
				/*case EaseType.bounce:
				return bounce(start, end, t);
				*/
			case EaseType.easeInBounce:
				return easeInBounce(start, end, t);
				
			case EaseType.easeOutBounce:
				return easeOutBounce(start, end, t);
				
			case EaseType.easeInOutBounce:
				return easeInOutBounce(start, end, t);
				
				/* GFX47 MOD END */
			case EaseType.easeInBack:
				return easeInBack(start, end, t);
				
			case EaseType.easeOutBack:
				return easeOutBack(start, end, t);
				
			case EaseType.easeInOutBack:
				return easeInOutBack(start, end, t);
				
				/* GFX47 MOD START */
				/*case EaseType.elastic:
				return elastic(start, end, t);
				*/
			case EaseType.easeInElastic:
				return easeInElastic(start, end, t);
				
			case EaseType.easeOutElastic:
				return easeOutElastic(start, end, t);
				
			case EaseType.easeInOutElastic:
				return easeInOutElastic(start, end, t);
				
				/* GFX47 MOD END */
			}
			return linear(start, end, t);
		}

		#endregion	

	}
}