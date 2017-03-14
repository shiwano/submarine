using System;

namespace TyphenApi
{
    public class MessagePackSerializer : ISerializer, IDeserializer
    {
        public byte[] Serialize<T>(T obj)
        {
            try
            {
                return MessagePack.MessagePackSerializer.Serialize(obj);
            }
            catch (Exception e)
            {
                throw new SerializationFailureException("Failed to serialize a object to a MessagePack data", e);
            }
        }

        public T Deserialize<T>(byte[] bytes) where T : new()
        {
            try
            {
                return MessagePack.MessagePackSerializer.Deserialize<T>(bytes);
            }
            catch (Exception e)
            {
                throw new DeserializationFailureException("Failed to deserialize a object from a MessagePack data", e);
            }
        }
    }
}
