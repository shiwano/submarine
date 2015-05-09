using System;

namespace Submarine
{
    /// <summary>
    /// Utility class for Unix time.
    /// Modified from https://gist.github.com/YuukiTsuchida/06ca3a1f0baf755651b0
    /// </summary>
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
