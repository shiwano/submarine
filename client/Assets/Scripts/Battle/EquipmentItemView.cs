using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Battle
{
    [RequireComponent(typeof(Button))]
    public class EquipmentItemView : MonoBehaviour
    {
        [SerializeField]
        Text cooldownText;

        Button button;

        public IObservable<Unit> OnUseAsObservable()
        {
            return button.OnSingleClickAsObservable();
        }

        void Awake()
        {
            button = GetComponent<Button>();
        }
    }
}
