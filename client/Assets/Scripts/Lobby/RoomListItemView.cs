using UnityEngine;
using UnityEngine.UI;

namespace Submarine.Lobby
{
    public class RoomListItemView : MonoBehaviour
    {
        [SerializeField]
        Button button;
        [SerializeField]
        Text roomIdText;
        [SerializeField]
        Text roomMembersText;
    }
}
