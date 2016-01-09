using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class TitleMediator : IInitializable
    {
        [Inject]
        PermanentDataStoreService dataStore;
        [Inject]
        LoginCommand loginCommand;
        [Inject]
        SignUpCommand signUpCommand;
        [Inject]
        SceneChangeCommand sceneChangeCommand;
        [Inject]
        DeleteLoginDataCommand deleteLoginDataCommand;
        [Inject]
        TitleEvents events;
        [Inject]
        TitleView view;

        public void Initialize()
        {
            events.LoginSucceeded.AddListener(OnLoginSuccess);
            view.StartButtonClickedAsObservable().Subscribe(_ => OnStartButtonClick());
            view.DeleteLoginButtonClickedAsObservable().Subscribe(_ => OnDeleteLoginDataButton());
        }

        void OnStartButtonClick()
        {
            if (dataStore.HasLoginData)
            {
                loginCommand.Execute();
            }
            else
            {
                signUpCommand.Execute("Test");
            }
        }

        void OnDeleteLoginDataButton()
        {
            deleteLoginDataCommand.Execute();
        }

        void OnLoginSuccess()
        {
            sceneChangeCommand.Execute(SceneNames.Battle);
        }
    }
}
