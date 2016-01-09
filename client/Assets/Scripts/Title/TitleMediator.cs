using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class TitleMediator : IInitializable
    {
        readonly PermanentDataStoreService dataStore;
        readonly LoginCommand loginCommand;
        readonly SignUpCommand signUpCommand;
        readonly SceneChangeCommand sceneChangeCommand;
        readonly DeleteLoginDataCommand deleteLoginDataCommand;
        readonly TitleEvents events;
        readonly TitleView view;

        public TitleMediator(
            PermanentDataStoreService dataStore,
            LoginCommand loginCommand,
            SignUpCommand signUpCommand,
            SceneChangeCommand sceneChangeCommand,
            DeleteLoginDataCommand deleteLoginDataCommand,
            TitleEvents events,
            TitleView view)
        {
            this.dataStore = dataStore;
            this.loginCommand = loginCommand;
            this.signUpCommand = signUpCommand;
            this.sceneChangeCommand = sceneChangeCommand;
            this.deleteLoginDataCommand = deleteLoginDataCommand;
            this.events = events;
            this.view = view;
        }

        public void Initialize()
        {
            events.LoginSucceeded.AddListener(OnLoginSuccess);
            view.StartButtonClickedAsObservable().Take(1).Subscribe(_ => OnStartButtonClick());
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
