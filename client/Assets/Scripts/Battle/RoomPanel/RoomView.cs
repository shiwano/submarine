using UnityEngine;
using UnityEngine.UI;
using System.Collections.Generic;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomView : MonoBehaviour, IView
    {
        [SerializeField]
        RectTransform roomMemberGroup;
        [SerializeField]
        RoomMemberView roomMemberTemplate;
        [SerializeField]
        Button battleStartButton;
        [SerializeField]
        Button addBotButton;

        Subject<Type.Bot> removeBotSubject;

        public IObservable<Unit> BattleStartButtonClickedAsObservable()
        {
            return battleStartButton.OnSingleClickAsObservable();
        }

        public IObservable<Unit> AddBotButtonClickedAsObservable()
        {
            return addBotButton.OnSingleClickAsObservable();
        }

        public IObservable<Type.Bot> RemoveBotButtonClickedAsObservable()
        {
            return removeBotSubject.AsObservable();
        }

        public void RefreshRoomMembers(IEnumerable<Type.User> users, IEnumerable<Type.Bot> bots)
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
                var roomMemberView = CreateRoomMemberView();
                roomMemberView.Setup(user);
            }

            if (bots != null)
            {
                foreach (var bot in bots)
                {
                    var roomMemberView = CreateRoomMemberView();
                    roomMemberView.Setup(bot);
                    roomMemberView.OnRemoveButtonClickAsObservable()
                        .Subscribe(v => removeBotSubject.OnNext(v.Bot));
                }
            }
        }

        void Awake()
        {
            removeBotSubject = new Subject<Type.Bot>();
            roomMemberTemplate.gameObject.SetActive(false);
        }

        RoomMemberView CreateRoomMemberView()
        {
            var roomMemberView = Instantiate<RoomMemberView>(roomMemberTemplate);
            roomMemberView.transform.SetParent(roomMemberGroup, false);
            roomMemberView.gameObject.SetActive(true);
            return roomMemberView;
        }
    }
}
