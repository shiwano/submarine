using TyphenApi.Type.Submarine;
using WebSocketSharp;
using UniRx;

namespace TyphenApi.WebSocketSession
{
    public class Submarine : Base.Submarine
    {
        public readonly ReactiveProperty<bool> IsConnected;

        public Submarine(string requestUri) : base(requestUri)
        {
            var jsonSerializer = new JSONSerializer();
            MessageSerializer = jsonSerializer;
            MessageDeserializer = jsonSerializer;

            IsConnected = new ReactiveProperty<bool>();
        }

        public override void OnConnectionCreate(WebSocket connection)
        {
        }

        public override void OnConnectionOpen()
        {
            IsConnected.Value = true;
        }

        public override void OnConnectionClose(ushort code, string reason, bool wasClean)
        {
            IsConnected.Value = false;
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
