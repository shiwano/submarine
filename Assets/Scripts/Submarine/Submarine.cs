using UnityEngine;
using DG.Tweening;

namespace Submarine
{
    public class Submarine
    {
        private readonly SubmarineHooks hooks;

        public Submarine(SubmarineHooks hooks)
        {
            this.hooks = hooks;
        }

        public void SetPositionAndRotation(Vector3 position, Vector3 eulerAngles)
        {
            hooks.transform.localPosition = position;
            hooks.transform.Rotate(eulerAngles);
        }

        public void MoveToVertically(float to, float duration, Ease easingType = Ease.Linear)
        {
            hooks.transform.DOMoveY(to, duration).SetEase(easingType);
        }
    }
}
