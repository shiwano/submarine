using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Lobby
{
    public class LobbyView : MonoBehaviour
    {
        [SerializeField]
        Button createRoomButton;

        public IObservable<Unit> CreateRoomButtonClickedAsObservable()
        {
            return createRoomButton.onClickAsObservableWithThrottle();
        }
    }
}
