using UnityEngine;
using UnityEngine.UI;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomMemberView : MonoBehaviour, IView
    {
        [SerializeField]
        Text nameText;
        [SerializeField]
        Button removeButton;

        public Type.User User { get; private set; }
        public Type.Bot Bot { get; private set; }

        public void Setup(Type.User user)
        {
            User = user;
            nameText.text = user.Name;
            removeButton.gameObject.SetActive(false);
        }

        public void Setup(Type.Bot bot)
        {
            Bot = bot;
            nameText.text = bot.Name;
            removeButton.gameObject.SetActive(true);
        }

        public IObservable<RoomMemberView> OnRemoveButtonClickAsObservable()
        {
            return removeButton.OnSingleClickAsObservable().Select(_ => this);
        }
    }
}
