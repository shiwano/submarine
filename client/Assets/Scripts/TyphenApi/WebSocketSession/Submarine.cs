using TyphenApi.Type.Submarine;
using WebSocketSharp;

namespace TyphenApi.WebSocketSession
{
    public class Submarine : Base.Submarine
    {
        public Submarine(string requestUri) : base(requestUri)
        {
            var jsonSerializer = new JSONSerializer();
            MessageSerializer = jsonSerializer;
            MessageDeserializer = jsonSerializer;
        }

        public override void OnConnectionCreate(WebSocket connection)
        {
        }

        public override void OnConnectionOpen()
        {
        }

        public override void OnConnectionClose(ushort code, string reason, bool wasClean)
        {
        }

        public override void OnBeforeMessageSend(TyphenApi.TypeBase message)
        {
        }

        public override void OnMessageReceive(TyphenApi.TypeBase message)
        {
        }

        public override void OnError(WebSocketSessionError<Error> error)
        {
        }
    }
}
