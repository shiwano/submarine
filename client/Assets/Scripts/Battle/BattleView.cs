using UnityEngine;
using UnityEngine.UI;
using System;
using System.Collections.Generic;

namespace Submarine.Battle
{
    public class BattleView : MonoBehaviour
    {
        [SerializeField]
        Text battleLogText;
        [SerializeField]
        Text timerText;

        [SerializeField]
        List<Image> torpedoResourceImages;
        [SerializeField]
        Image pingerAlertImage;
        [SerializeField]
        Image dangerAlertImage;

        [SerializeField]
        Text decoyCoolDownText;
        [SerializeField]
        Text pingerCoolDownText;
        [SerializeField]
        Text lookoutCoolDownText;

        public bool IsUsingPinger
        {
            set { pingerAlertImage.gameObject.SetActive(value); }
        }

        public TimeSpan ElapsedTime
        {
            set { timerText.text = string.Format("{0:00}:{1:00}", (int)value.TotalMinutes, (int)value.Seconds); }
        }
    }
}
