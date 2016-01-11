using UnityEngine;
using UnityEngine.UI;
using System;
using System.Collections.Generic;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Lobby
{
    public class RoomListView : MonoBehaviour
    {
        [SerializeField]
        RectTransform content;
        [SerializeField]
        RoomListItem listItemTemplate;
        [SerializeField]
        Button createRoomButton;
        [SerializeField]
        Button updateRoomsButton;

        public IObservable<Unit> CreateRoomButtonClickedAsObservable()
        {
            return createRoomButton.onClickAsObservableWithThrottle();
        }

        public IObservable<Unit> UpdateRoomsButtonClickedAsObservable()
        {
            return updateRoomsButton.onClickAsObservableWithThrottle();
        }

        public void ClearRooms()
        {
            foreach (Transform listItem in content.transform)
            {
                Destroy(listItem.gameObject);
            }
        }

        public void CreateRooms(IEnumerable<Type.Room> rooms, Action<Type.Room> onClick)
        {
            foreach (var room in rooms)
            {
                var listItem = Instantiate<RoomListItem>(listItemTemplate);
                listItem.transform.SetParent(content, false);
                listItem.Setup(room, onClick);
            }
        }

        void Awake()
        {
            listItemTemplate.gameObject.SetActive(false);
        }
    }
}
