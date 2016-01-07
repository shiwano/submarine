using UniRx;
using Zenject;

namespace Submarine
{
    public class TitleMediator : IInitializable
    {
        readonly Services.PermanentDataStore dataStore;
        readonly Commands.Login loginCommand;
        readonly Commands.SignUp signUpCommand;
        readonly Commands.SceneChange sceneChangeCommand;
        readonly Events.Title events;
        readonly TitleView view;

        public TitleMediator(
            Services.PermanentDataStore dataStore,
            Commands.Login loginCommand,
            Commands.SignUp signUpCommand,
            Commands.SceneChange sceneChangeCommand,
            Events.Title events,
            TitleView view)
        {
            this.dataStore = dataStore;
            this.loginCommand = loginCommand;
            this.signUpCommand = signUpCommand;
            this.sceneChangeCommand = sceneChangeCommand;
            this.events = events;
            this.view = view;
        }

        public void Initialize()
        {
            events.LoginSucceeded.AddListener(OnLoginSuccess);
            view.StartButtonClickedAsObservable().Take(1).Subscribe(_ => OnStartButtonClick());
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

        void OnLoginSuccess()
        {
            sceneChangeCommand.Execute(SceneNames.Battle);
        }
    }
}
