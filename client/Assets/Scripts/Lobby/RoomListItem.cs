using UnityEngine;
using UnityEngine.UI;
using System.Linq;
using System;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class RoomListItem : MonoBehaviour
    {
        [SerializeField]
        Button button;
        [SerializeField]
        Text roomIdText;
        [SerializeField]
        Text roomMembersText;

        public void Setup(Type.Room room, Action<Type.Room> onClick)
        {
            roomIdText.text = room.Id.ToString();
            roomMembersText.text = room.Members.Select(m => m.Name).Aggregate((a, b) => a + ", " + b);
            button.OnSingleClickAsObservable().Subscribe(_ => onClick(room)).AddTo(this);
        }
    }
}
