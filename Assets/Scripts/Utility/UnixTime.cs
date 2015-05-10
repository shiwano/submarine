using System;

namespace Submarine
{
    public static class UnixTime
    {
        static readonly DateTime UnixEpoch = new DateTime(1970, 1, 1, 0, 0, 0, DateTimeKind.Utc);

        public static long Now
        {
            get { return FromDateTime(DateTime.UtcNow); }
        }

        public static DateTime FromUnixTime(long unixTime)
        {
            return UnixEpoch.AddSeconds(unixTime).ToLocalTime();
        }

        public static long FromDateTime(DateTime dateTime)
        {
            double nowTicks = (dateTime.ToUniversalTime() - UnixEpoch).TotalSeconds;
            return (long)nowTicks;
        }
    }
}
