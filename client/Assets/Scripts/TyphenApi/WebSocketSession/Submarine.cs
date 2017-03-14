using UnityEngine;
using TyphenApi.Type.Submarine;
using WebSocketSharp;
using UniRx;
using Game = Submarine;

namespace TyphenApi.WebSocketSession
{
    public class Submarine : Base.Submarine
    {
        CompositeDisposable disposables = new CompositeDisposable();

        public readonly ReactiveProperty<bool> IsConnected;

        public Submarine(string requestUri) : base(requestUri)
        {
            var jsonSerializer = new JSONSerializer();
            MessageSerializer = jsonSerializer;
            MessageDeserializer = jsonSerializer;

            IsConnected = new ReactiveProperty<bool>();
            Observable.EveryUpdate().Subscribe(_ => Update()).AddTo(disposables);
        }

        public override void OnConnectionCreate(WebSocket connection)
        {
        }

        public override void OnConnectionOpen()
        {
            IsConnected.Value = true;
            Game.Logger.Log("[WebSocketApi] Session opened");
        }

        public override void OnConnectionClose(ushort code, string reason, bool wasClean)
        {
            IsConnected.Value = false;
            disposables.Dispose();
            disposables = new CompositeDisposable();
            Game.Logger.Log("[WebSocketApi] Session closed");
        }

        public override void OnBeforeMessageSend(TyphenApi.IType message)
        {
            #if UNITY_EDITOR
            Game.Logger.LogWithBlue("[WebSocketApi] Send " + message.GetType().Name + " Message:", message);
            #endif
        }

        public override void OnMessageReceive(TyphenApi.IType message)
        {
            #if UNITY_EDITOR
            Game.Logger.LogWithGreen("[WebSocketApi] Receive " + message.GetType().Name + " Message:", message);
            #endif
        }

        public override void OnError(WebSocketSessionError<Error> error)
        {
            #if UNITY_EDITOR
            if (error.Error != null)
            {
                Game.Logger.LogError("[WebSocketApi] Error:", error.Error.Code + ": " + error.Error.Name, error.Error);
            }
            else
            {
                Game.Logger.LogError("[WebSocketApi] Error:", error.RawErrorMessage);
            }
            #endif
        }
    }
}

