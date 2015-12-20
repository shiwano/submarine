using ModestTree.Util;
using System;

namespace Zenject
{
    [AttributeUsage(AttributeTargets.Parameter | AttributeTargets.Property | AttributeTargets.Field, AllowMultiple = false)]
    public abstract class InjectAttributeBase : PreserveAttribute
    {
        public bool IsOptional
        {
            get;
            protected set;
        }

        public string Identifier
        {
            get;
            protected set;
        }

        public bool LocalOnly
        {
            get;
            protected set;
        }
    }
}


