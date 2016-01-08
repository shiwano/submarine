using UnityEngine;
using System.Linq;

namespace Submarine
{
    public static class Logger
    {
        public static void Log(params object[] messages)
        {
            UnityEngine.Debug.Log(BuildMessage(messages));
        }

        public static void LogWithColor(Color color, params object[] messages)
        {
            UnityEngine.Debug.Log(BuildMessage(messages, color));
        }

        public static void LogError(params object[] messages)
        {
            UnityEngine.Debug.LogError(BuildMessage(messages, Color.red));
        }

        static string BuildMessage(object[] messages, Color? color = null)
        {
            var message = string.Join("\n", messages.Select(m => m.ToString()).ToArray());
            if (color != null)
            {
                message = Colorize(message, color.Value);
            }
            return message;
        }

        static string Colorize(string str, Color color)
        {
            var r = (int)(color.r * 255);
            var g = (int)(color.g * 255);
            var b = (int)(color.b * 255);
            var hex = r.ToString("X2") + g.ToString("X2") + b.ToString("X2");
            return string.Format ("<color=#{0}>{1}</color>", hex, str);
        }
    }
}