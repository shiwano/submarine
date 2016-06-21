using UnityEngine;
using UnityEngine.UI;
using System.Collections.Generic;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomView : MonoBehaviour
    {
        [SerializeField]
        RectTransform roomMemberGroup;
        [SerializeField]
        RoomMemberView roomMemberTemplate;
        [SerializeField]
        Button battleStartButton;

        public IObservable<Unit> BattleStartButtonClickedAsObservable()
        {
            return battleStartButton.OnSingleClickAsObservable();
        }

        public void RefreshRoomMembers(IEnumerable<Type.User> users)
        {
            foreach (Transform roomMember in roomMemberGroup.transform)
            {
                if (roomMember.gameObject != roomMemberTemplate.gameObject)
                {
                    Destroy(roomMember.gameObject);
                }
            }

            foreach (var user in users)
            {
                var listItem = Instantiate<RoomMemberView>(roomMemberTemplate);
                listItem.Setup(user);
                listItem.transform.SetParent(roomMemberGroup, false);
                listItem.gameObject.SetActive(true);
            }
        }

        void Awake()
        {
            roomMemberTemplate.gameObject.SetActive(false);
        }
    }
}
