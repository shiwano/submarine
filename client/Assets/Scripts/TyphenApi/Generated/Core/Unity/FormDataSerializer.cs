using UnityEngine;
using System;

namespace TyphenApi
{
    public class FormDataSerializer : ISerializer
    {
        public byte[] Serialize(object obj)
        {
            var objType = obj.GetType();
            var form = new WWWForm();

            foreach (var info in SerializationInfoFinder.FindAll(obj))
            {
                var value = info.GetValue(obj);
                var valueType = info.ValueType;

                if (value == null)
                {
                    if (info.IsOptional)
                    {
                        continue;
                    }
                    else
                    {
                        var message = string.Format("{0}.{1} is not allowed to be null.", objType.FullName, info.PropertyName);
                        throw new NoNullAllowedException(message);
                    }
                }

                if (IsSerializableValue(value, valueType))
                {
                    var fixedValue = valueType.IsEnum ? (int)value : value;
                    form.AddField(info.PropertyName, fixedValue.ToString());
                }
                else
                {
                    var message = string.Format("Failed to serialize {0} ({1}) to {2}.{3}", value, valueType, objType.FullName, info.PropertyName);
                    throw new SerializeFailedException(message);
                }
            }
            return form.data;
        }

        bool IsSerializableValue(object value, System.Type valueType)
        {
            return valueType.IsPrimitive || valueType.IsEnum || value is string;
        }
    }
}
