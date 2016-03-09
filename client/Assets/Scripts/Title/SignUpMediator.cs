using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class SignUpMediator : IInitializable
    {
        [Inject]
        TitleEvent.SignUpStart signUpStartEvent;
        [Inject]
        SignUpCommand signUpCommand;
        [Inject]
        SignUpView view;

        public void Initialize()
        {
            signUpStartEvent.AsObservable().Subscribe(_ => OnSignUpStart()).AddTo(view);
            view.SignUpButtonClickedAsObservable().Subscribe(_ => OnSignUpButtonClick()).AddTo(view);
            view.Hide();
        }

        public void OnSignUpStart()
        {
            view.Show();
            view.FocusToNameInputField();
        }

        public void OnSignUpButtonClick()
        {
            if (!string.IsNullOrEmpty(view.InputtedName))
            {
                signUpCommand.Execute(view.InputtedName);
            }
        }
    }
}
