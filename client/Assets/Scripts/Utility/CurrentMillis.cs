using System;

namespace Submarine
{
    public static class CurrentMillis
    {
        static readonly DateTime UnixEpoch = new DateTime(1970, 1, 1, 0, 0, 0, DateTimeKind.Utc);

        public static long Now
        {
            get { return FromDateTime(DateTime.UtcNow); }
        }

        public static DateTime FromMilliseconds(long ms)
        {
            return CurrentMillis.UnixEpoch.AddMilliseconds(ms).ToLocalTime();
        }

        public static long FromDateTime(DateTime dateTime)
        {
            return (long)(dateTime.ToUniversalTime() - CurrentMillis.UnixEpoch).TotalMilliseconds;
        }
    }
}
