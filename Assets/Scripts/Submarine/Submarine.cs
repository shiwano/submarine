using UnityEngine;
using System.Collections;
using Zenject;
using DG.Tweening;

namespace Submarine
{
    public class Submarine
    {
        private readonly SubmarineView view;

        public Submarine(SubmarineView view)
        {
            this.view = view;
        }

        public void SetPositionAndRotation(Vector3 position, Vector3 eulerAngles)
        {
            view.transform.localPosition = position;
            view.transform.Rotate(eulerAngles);
        }

        public void MoveToVertically(float to, float duration, Ease easingType = Ease.Linear)
        {
            view.transform.DOMoveY(to, duration).SetEase(easingType);
        }
    }
}
