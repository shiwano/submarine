using System;

namespace Zenject
{
    [AttributeUsage(AttributeTargets.Constructor | AttributeTargets.Parameter | AttributeTargets.Property | AttributeTargets.Field, AllowMultiple = false)]
    public class InjectAttribute : InjectAttributeBase
    {
        public InjectAttribute(string identifier)
        {
            Identifier = identifier;
        }

        public InjectAttribute()
        {
        }
    }
}

