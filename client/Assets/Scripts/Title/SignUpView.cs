using UnityEngine;
using UnityEngine.UI;
using UniRx;

namespace Submarine.Title
{
    public class SignUpView : MonoBehaviour
    {
        [SerializeField]
        Button signUpButton;
        [SerializeField]
        InputField nameInputField;

        public string InputtedName { get { return nameInputField.text; } }

        public IObservable<Unit> SignUpButtonClickedAsObservable()
        {
            return signUpButton.onClickAsObservableWithThrottle();
        }

        public IObservable<string> NameChangedAsObservable()
        {
            return nameInputField.OnValueChangeAsObservable();
        }

        public void Show()
        {
            gameObject.SetActive(true);
        }

        public void Hide()
        {
            gameObject.SetActive(false);
        }

        public void FocusToNameInputField()
        {
            nameInputField.ActivateInputField();
            nameInputField.Select();
        }
    }
}
