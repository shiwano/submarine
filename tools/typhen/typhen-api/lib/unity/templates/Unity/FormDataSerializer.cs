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

            foreach (var property in objType.GetProperties())
            {
                var attributes = property.GetCustomAttributes(typeof(SerializablePropertyAttribute), true);

                foreach (var attribute in attributes)
                {
                    var metaInfo = (SerializablePropertyAttribute)attribute;
                    var value = property.GetValue(obj, null);

                    if (metaInfo.IsOptional && value == null)
                    {
                        continue;
                    }

                    var valueType = property.PropertyType;

                    if (value == null && IsNullableType(valueType))
                    {
                        var message = string.Format("{0}.{1} is not allowed to be null.", objType.FullName, property.Name);
                        throw new NoNullAllowedException(message);
                    }
                    else if (IsSerializableValue(value, valueType))
                    {
                        var fixedValue = valueType.IsEnum ? (int)value : value;
                        form.AddField(metaInfo.PropertyName, fixedValue.ToString());
                    }
                    else
                    {
                        var message = string.Format("Failed to serialize {0} ({1}) to {2}.{3}", value, valueType, objType.FullName, property.Name);
                        throw new SerializeFailedException(message);
                    }
                }
            }

            return form.data;
        }

        bool IsNullableType(System.Type type)
        {
            return type.IsClass || Nullable.GetUnderlyingType(type) != null;
        }

        bool IsSerializableValue(object value, System.Type valueType)
        {
            return valueType.IsPrimitive || valueType.IsEnum || value is string;
        }
    }
}
