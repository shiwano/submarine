using UnityEngine;
using UnityEngine.UI;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class RoomMemberView : MonoBehaviour
    {
        [SerializeField]
        Text nameText;

        public void Setup(Type.User user)
        {
            nameText.text = user.Name;
        }
    }
}
