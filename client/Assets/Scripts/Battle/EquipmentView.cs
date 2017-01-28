using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Battle
{
    public class EquipmentView : MonoBehaviour, BattleInputService.IEquipmentInput
    {
        [SerializeField]
        EquipmentItemView decoy;
        [SerializeField]
        EquipmentItemView pinger;
        [SerializeField]
        EquipmentItemView watcher;

        [SerializeField]
        Image torpedoResourceImageTemplate;
        [SerializeField]
        RectTransform torpedoResourceImageRoot;
        [SerializeField]
        Vector3 torpedoResourceImageOffset;

        public IObservable<Unit> OnDecoyUseAsObservable()   { return decoy.OnUseAsObservable(); }
        public IObservable<Unit> OnPingerUseAsObservable()  { return pinger.OnUseAsObservable(); }
        public IObservable<Unit> OnWatcherUseAsObservable() { return watcher.OnUseAsObservable(); }
    }
}
