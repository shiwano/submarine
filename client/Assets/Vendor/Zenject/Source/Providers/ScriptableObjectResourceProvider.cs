#if !NOT_UNITY3D

using System;
using System.Collections.Generic;
using System.Linq;
using ModestTree;
using UnityEngine;

namespace Zenject
{
    public class ScriptableObjectResourceProvider : IProvider
    {
        readonly DiContainer _container;
        readonly Type _resourceType;
        readonly string _resourcePath;
        readonly List<TypeValuePair> _extraArguments;
        readonly object _concreteIdentifier;

        public ScriptableObjectResourceProvider(
            string resourcePath, Type resourceType,
            DiContainer container, object concreteIdentifier, List<TypeValuePair> extraArguments)
        {
            _container = container;
            Assert.DerivesFromOrEqual<ScriptableObject>(resourceType);

            _concreteIdentifier = concreteIdentifier;
            _extraArguments = extraArguments;
            _resourceType = resourceType;
            _resourcePath = resourcePath;
        }

        public Type GetInstanceType(InjectContext context)
        {
            return _resourceType;
        }

        public IEnumerator<List<object>> GetAllInstancesWithInjectSplit(
            InjectContext context, List<TypeValuePair> args)
        {
            Assert.IsEmpty(args);

            Assert.IsNotNull(context);

            var objects = Resources.LoadAll(_resourcePath, _resourceType)
                .Select(x => ScriptableObject.Instantiate(x)).Cast<object>().ToList();

            Assert.That(!objects.IsEmpty(),
                "Could not find resource at path '{0}' with type '{1}'", _resourcePath, _resourceType);

            yield return objects;

            var injectArgs = new InjectArgs()
            {
                ExtraArgs = _extraArguments.Concat(args).ToList(),
                Context = context,
                ConcreteIdentifier = _concreteIdentifier,
            };

            foreach (var obj in objects)
            {
                _container.InjectExplicit(
                    obj, _resourceType, injectArgs);
            }
        }
    }
}

#endif
