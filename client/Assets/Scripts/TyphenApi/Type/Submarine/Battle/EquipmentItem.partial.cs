using System;
using Submarine;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class EquipmentItem
    {
        public DateTime CooldownStartedAtAsDateTime
        {
            get { return CurrentMillis.FromMilliseconds(CooldownStartedAt); }
        }

        public DateTime CooldownFinishedAtAsDateTime
        {
            get { return CurrentMillis.FromMilliseconds(CooldownStartedAt + CooldownDuration); }
        }
    }
}
