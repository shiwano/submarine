using System;
using System.Linq;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public enum BattleState
    {
        NotConnected,
        InPreparation,
        InBattle,
        Finish,
    }

    public class BattleModel
    {
        public readonly ReactiveProperty<BattleState> State = new ReactiveProperty<BattleState>(BattleState.NotConnected);
        public readonly ReactiveDictionary<long, Type.Battle.Actor> ActorsById = new ReactiveDictionary<long, Type.Battle.Actor>();

        public DateTime StartedAt { get; set; }
        public DateTime FinishedAt { get; set; }
        public Type.User Winner { get; set; }

        public TimeSpan ElapsedTime { get { return Now - StartedAt; } }
        public bool IsInBattle { get { return State.Value == BattleState.InBattle; } }

        TimeSpan differenceFromBattleServerTime;
        public DateTime Now
        {
            get { return DateTime.Now.Add(differenceFromBattleServerTime); }
            set { differenceFromBattleServerTime = value.Subtract(DateTime.Now); }
        }

        public IObservable<Unit> OnPrepareAsObservable()
        {
            return State.Where(s => s == BattleState.InPreparation).AsUnitObservable();
        }

        public IObservable<Unit> OnStartAsObservable()
        {
            return State.Where(s => s == BattleState.InBattle).AsUnitObservable();
        }

        public IObservable<Unit> OnFinishAsObservable()
        {
            return State.Where(s => s == BattleState.Finish).AsUnitObservable();
        }
    }
}
