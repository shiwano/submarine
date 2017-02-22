using System;
using UnityEngine;
using UnityEngine.UI;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    [RequireComponent(typeof(Button))]
    public class EquipmentItemView : MonoBehaviour, IView
    {
        [SerializeField]
        Text cooldownText;

        Button button;

        public IObservable<Unit> OnUseAsObservable()
        {
            return button.OnSingleClickAsObservable();
        }

        public void Refresh(DateTime now, Type.Battle.EquipmentItem equipmentItem)
        {
            var finishedAt = equipmentItem.CooldownFinishedAtAsDateTime;
            button.interactable = now >= finishedAt;
            cooldownText.text = button.interactable ?
                string.Empty :
                string.Format("{0:00}", (finishedAt - now).TotalSeconds);
        }

        void Awake()
        {
            button = GetComponent<Button>();
            button.interactable = false;
        }
    }
}
