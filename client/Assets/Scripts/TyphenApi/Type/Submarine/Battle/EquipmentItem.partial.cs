using System;
using Submarine;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class EquipmentItem
    {
        [MessagePack.IgnoreMember]
        public DateTime CooldownStartedAtAsDateTime
        {
            get { return CurrentMillis.FromMilliseconds(CooldownStartedAt); }
        }

        [MessagePack.IgnoreMember]
        public DateTime CooldownFinishedAtAsDateTime
        {
            get { return CurrentMillis.FromMilliseconds(CooldownStartedAt + CooldownDuration); }
        }
    }
}
