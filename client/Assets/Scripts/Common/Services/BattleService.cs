using UniRx;

namespace Submarine
{
    public class BattleService
    {
        TyphenApi.WebSocketSession.Submarine session;

        public ReactiveProperty<bool> IsStarted { get; private set; }
        public TyphenApi.WebSocketApi.Parts.Submarine.Battle Api { get { return session.Api.Battle; } }

        public BattleService()
        {
            IsStarted = new ReactiveProperty<bool>();
        }

        public void Start(string requestUri)
        {
            session = new TyphenApi.WebSocketSession.Submarine(requestUri);
            session.IsConnected.Skip(1).Take(2).Subscribe(
                v => IsStarted.Value = v,
                () => session = null);
            session.Connect();
        }

        public void Stop()
        {
            if (session != null && session.IsOpened)
            {
                session.Dispose();
            }
        }
    }
}
