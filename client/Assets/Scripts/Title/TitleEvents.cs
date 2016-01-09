using UnityEngine.Events;

namespace Submarine.Title
{
    public class TitleEvents
    {
        public readonly UnityEvent LoginSucceeded = new UnityEvent();
        public readonly UnityEvent SignUpStarted = new UnityEvent();
    }
}
