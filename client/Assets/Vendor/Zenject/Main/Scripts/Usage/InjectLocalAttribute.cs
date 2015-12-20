using System;

namespace Zenject
{
    [AttributeUsage(AttributeTargets.Parameter | AttributeTargets.Property | AttributeTargets.Field, AllowMultiple = false)]
    public class InjectLocalAttribute : InjectAttributeBase
    {
        public InjectLocalAttribute(string identifier)
        {
            Identifier = identifier;
            LocalOnly = true;
        }

        public InjectLocalAttribute()
        {
            LocalOnly = true;
        }
    }

    [AttributeUsage(AttributeTargets.Parameter | AttributeTargets.Property | AttributeTargets.Field, AllowMultiple = false)]
    public class InjectLocalOptionalAttribute : InjectAttributeBase
    {
        public InjectLocalOptionalAttribute(string identifier)
        {
            Identifier = identifier;
            IsOptional = true;
            LocalOnly = true;
        }

        public InjectLocalOptionalAttribute()
        {
            IsOptional = true;
            LocalOnly = true;
        }
    }
}

