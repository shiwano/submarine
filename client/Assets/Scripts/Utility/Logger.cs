using UnityEngine;
using System.Linq;

namespace Submarine
{
    public static class Logger
    {
        public static void Log(params object[] messages)
        {
            Debug.Log(BuildMessage(messages));
        }

        public static void LogWithColor(Color color, params object[] messages)
        {
            Debug.Log(BuildMessage(messages, color));
        }

        public static void LogError(params object[] messages)
        {
            Debug.LogError(BuildMessage(messages, Color.red));
        }

        static string BuildMessage(object[] messageObjs, Color? color = null)
        {
            var messages = messageObjs
                .Where(m => m != null)
                .Select(messageObj =>
                {
                    return color != null ?
                        Colorize(messageObj.ToString(), color.Value) :
                        messageObj.ToString();
                })
                .ToArray();
            return string.Join("\n", messages);
        }

        static string Colorize(string str, Color color)
        {
            var r = (int)(color.r * 255);
            var g = (int)(color.g * 255);
            var b = (int)(color.b * 255);
            var hex = r.ToString("X2") + g.ToString("X2") + b.ToString("X2");
            return string.Format("<color=#{0}>{1}</color>", hex, str);
        }
    }
}