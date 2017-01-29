using System;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using UniRx;
using Type = TyphenApi.Type.Submarine;

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

        readonly List<Image> torpedoResourceImages = new List<Image>();

        public IObservable<Unit> OnDecoyUseAsObservable()   { return decoy.OnUseAsObservable(); }
        public IObservable<Unit> OnPingerUseAsObservable()  { return pinger.OnUseAsObservable(); }
        public IObservable<Unit> OnWatcherUseAsObservable() { return watcher.OnUseAsObservable(); }

        public void Refresh(DateTime now, Type.Battle.Equipment equipment)
        {
            RefreshTorpedos(now, equipment);
            pinger.Refresh(now, equipment.Pinger);
        }

        void Awake()
        {
            torpedoResourceImageTemplate.gameObject.SetActive(false);
        }

        void RefreshTorpedos(DateTime now, Type.Battle.Equipment equipment)
        {
            RefreshTorpedoResourceImagesIfNeeded(equipment);

            equipment.Torpedos.ForEach((torpedo, i) =>
            {
                var color = Color.white;
                color.a = now >= torpedo.CooldownFinishedAtAsDateTime ? 1f : 0.35f;
                torpedoResourceImages[i].color = color;
            });
        }

        void RefreshTorpedoResourceImagesIfNeeded(Type.Battle.Equipment equipment)
        {
            if (torpedoResourceImages.Count == equipment.Torpedos.Count) return;

            torpedoResourceImages.ForEach(i => Destroy(i.gameObject));
            torpedoResourceImages.Clear();

            for (var i = 0; i < equipment.Torpedos.Count; i++)
            {
                var image = Instantiate(torpedoResourceImageTemplate).GetComponent<Image>();
                image.rectTransform.SetParent(torpedoResourceImageRoot, false);
                image.rectTransform.localPosition += torpedoResourceImageOffset * i;
                image.gameObject.SetActive(true);
                torpedoResourceImages.Add(image);
            }
        }
    }
}
