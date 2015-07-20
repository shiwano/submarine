using System;

namespace TyphenApi
{
    public abstract class SerializationError : Exception
    {
        protected SerializationError(string message) : base(message) {}
    }

    public class NoNullAllowedException : SerializationError
    {
        public NoNullAllowedException(string message) : base(message) {}
    }

    public class SerializeFailedError : SerializationError
    {
        public SerializeFailedError(string message) : base(message) {}
    }

    public class DeserializeFailedError : SerializationError
    {
        public DeserializeFailedError(string message) : base(message) {}
    }
}
