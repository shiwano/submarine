using UnityEngine;
using UnityEngine.UI;
using UnityEngine.EventSystems;
using System.Collections;

namespace uTools {
	[AddComponentMenu("uTools/Internal/Play Tween(uTools)")]
	public class uPlayTween : MonoBehaviour, uIPointHandler {
		public uTweener tweenTarget;
		public PlayDirection playDirection = PlayDirection.Forward;
		public Trigger trigger = Trigger.OnPointerClick;

		// Use this for initialization
		void Start () {
			if (tweenTarget == null) {
				tweenTarget = GetComponent<uTweener>();
			}		
		}

		public void OnPointerEnter (PointerEventData eventData) {
			TriggerPlay (Trigger.OnPointerEnter);
		}

		public void OnPointerDown (PointerEventData eventData) {
			TriggerPlay (Trigger.OnPointerDown);
		}

		public void OnPointerClick (PointerEventData eventData) {
			TriggerPlay (Trigger.OnPointerClick);
		}

		public void OnPointerUp (PointerEventData eventData) {
			TriggerPlay (Trigger.OnPointerUp);
		}

		public void OnPointerExit (PointerEventData eventData) {
			TriggerPlay (Trigger.OnPointerExit);
		}

		private void TriggerPlay(Trigger _trigger) {
			if (_trigger == trigger) {
				Play();
			}
		}

		/// <summary>
		/// Play this instance.
		/// </summary>
		private void Play() {
			if (playDirection == PlayDirection.Toggle) {
				tweenTarget.Toggle();
			}
			else {
				tweenTarget.Play(playDirection);
			}
		}

	}
}