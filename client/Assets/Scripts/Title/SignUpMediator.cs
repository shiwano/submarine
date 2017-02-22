using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class SignUpMediator : MediatorBase<SignUpView>, IInitializable
    {
        [Inject]
        TitleEvent.SignUpStart signUpStartEvent;
        [Inject]
        SignUpCommand signUpCommand;

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
                signUpCommand.Fire(view.InputtedName);
            }
        }
    }
}
