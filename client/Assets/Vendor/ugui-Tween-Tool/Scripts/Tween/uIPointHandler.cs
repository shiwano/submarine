using UnityEngine;
using UnityEngine.EventSystems;
using System.Collections;

namespace uTools {
	public interface uIPointHandler : 
		IPointerEnterHandler,
		IPointerDownHandler,
		IPointerClickHandler,
		IPointerUpHandler,
		IPointerExitHandler {

		new void OnPointerEnter (PointerEventData eventData);
		new void OnPointerDown (PointerEventData eventData);
		new void OnPointerClick (PointerEventData eventData);
		new void OnPointerUp (PointerEventData eventData);
		new void OnPointerExit (PointerEventData eventData);

	}

}