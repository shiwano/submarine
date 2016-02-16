using UnityEngine;
using UnityEngine.UI;
using System;
using System.Collections.Generic;

namespace Submarine.Battle
{
    public class BattleView : MonoBehaviour
    {
        public Camera MainCamera;

        public Text BattleLogText;
        public Text TimerText;

        public List<Image> TorpedoResourceImages;
        public Image PingerAlert;
        public Image DangerAlert;

        public Button Victory;
        public Button Defeat;

        public Text DecoyCoolDown;
        public Text PingerCoolDown;
        public Text LookoutCoolDown;

        public TimeSpan ElapsedTime
        {
            set
            {
                TimerText.text = string.Format(
                    "{0:00}:{1:00}",
                    (int)value.TotalMinutes,
                    (int)value.Seconds
                );
            }
        }
    }
}
