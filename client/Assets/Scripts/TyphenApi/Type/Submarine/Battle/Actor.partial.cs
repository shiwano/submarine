using UniRx;

namespace TyphenApi.Type.Submarine.Battle
{
    public partial class Actor
    {
        readonly Subject<Movement> onMoveSubject = new Subject<Movement>();
        readonly Subject<Equipment> onChangeEquipmentSubject = new Subject<Equipment>();

        public void UpdateValues(Movement movement)
        {
            Movement = movement;
            onMoveSubject.OnNext(Movement);
        }

        public void UpdateValues(Visibility visibility)
        {
            IsVisible = visibility.IsVisible;
            Movement = visibility.Movement;
            onMoveSubject.OnNext(Movement);
        }

        public void UpdateValues(Pinger pinger)
        {
            if (Submarine != null)
            {
                Submarine.IsUsingPinger = !pinger.IsFinished;
            }
        }

        public void UpdateValues(Equipment equipment)
        {
            if (Submarine != null)
            {
                Submarine.Equipment = equipment;
                onChangeEquipmentSubject.OnNext(equipment);
            }
        }

        public IObservable<Movement> OnMoveAsObservable()             { return onMoveSubject; }
        public IObservable<Equipment> OnChangeEquipmentAsObservable() { return onChangeEquipmentSubject; }
    }
}
