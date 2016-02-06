using System;
using UniRx;

namespace Submarine
{
    public class BattleService : IDisposable
    {
        TyphenApi.WebSocketSession.Submarine session;

        public ReactiveProperty<bool> IsConnected { get; private set; }
        public TyphenApi.WebSocketApi.Parts.Submarine.Battle Api { get { return session.Api.Battle; } }

        public BattleService()
        {
            IsConnected = new ReactiveProperty<bool>();
        }

        public void Connect(string requestUri)
        {
            session = new TyphenApi.WebSocketSession.Submarine(requestUri);
            session.IsConnected.Skip(1).Take(2).Subscribe(
                v => IsConnected.Value = v,
                () => session = null);
            session.Connect();
        }

        public void Dispose()
        {
            if (session != null && session.IsOpened)
            {
                session.Dispose();
            }
        }
    }
}
