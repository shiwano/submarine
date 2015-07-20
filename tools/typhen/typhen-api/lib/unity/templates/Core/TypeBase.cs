using System;
using System.Text;

namespace TyphenApi
{
    public abstract class TypeBase
    {
        public override string ToString()
        {
            var bytes = new JSONSerializer().Serialize(this);
            return Encoding.UTF8.GetString(bytes);
        }

        public string ToQueryString()
        {
            var bytes = new QueryStringSerializer().Serialize(this);
            return Encoding.UTF8.GetString(bytes);
        }
    }
}
