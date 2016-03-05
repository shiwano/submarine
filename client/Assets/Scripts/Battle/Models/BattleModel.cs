using System;
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
        public readonly ReactiveProperty<BattleState> State;

        public DateTime StartedAt { get; set; }
        public DateTime FinishedAt { get; set; }

        TimeSpan differenceFromBattleServerTime;
        public DateTime Now
        {
            get { return DateTime.Now.Add(differenceFromBattleServerTime); }
            set { differenceFromBattleServerTime = value.Subtract(DateTime.Now); }
        }

        public TimeSpan ElapsedTime
        {
            get { return Now - StartedAt; }
        }

        public bool IsInBattle
        {
            get { return State.Value == BattleState.InBattle; }
        }

        public BattleModel()
        {
            State = new ReactiveProperty<BattleState>(BattleState.InPreparation);
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
