using System;
using System.Collections.Generic;
 
namespace Submarine
{
    public static class IEnumerableExtensions
    {
        public static void ForEach<T>(this IEnumerable<T> enumerable, Action<T> callback)
        {
            foreach (var item in enumerable)
            {
                callback(item);
            }
        }

        public static void ForEach<T>(this IEnumerable<T> enumerable, Action<T, int> callback)
        {
            var index = 0;

            foreach (var item in enumerable)
            {
                callback(item, index);
                index++;
            }
        }
    }
}