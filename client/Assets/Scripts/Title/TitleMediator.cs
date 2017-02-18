using UniRx;
using Zenject;

namespace Submarine.Title
{
    public class TitleMediator : IInitializable
    {
        [Inject]
        TitleEvent.SignUpStart signUpStartEvent;
        [Inject]
        UserModel userModel;
        [Inject]
        PermanentDataStoreService dataStore;
        [Inject]
        LoginCommand loginCommand;
        [Inject]
        SceneChangeCommand sceneChangeCommand;
        [Inject]
        DeleteLoginDataCommand deleteLoginDataCommand;
        [Inject]
        TitleView view;

        public void Initialize()
        {
            userModel.HasLoggedIn.Where(v => v).Subscribe(_ => OnUserLogin()).AddTo(view);
            view.StartButtonClickedAsObservable().Subscribe(_ => OnStartButtonClick()).AddTo(view);
            view.DeleteLoginButtonClickedAsObservable().Subscribe(_ => OnDeleteLoginDataButton()).AddTo(view);
        }

        void OnStartButtonClick()
        {
            if (dataStore.HasSignedUp)
            {
                loginCommand.Fire();
            }
            else
            {
                signUpStartEvent.Invoke();
            }
        }

        void OnDeleteLoginDataButton()
        {
            deleteLoginDataCommand.Fire();
        }

        void OnUserLogin()
        {
            sceneChangeCommand.Fire(SceneNames.Lobby);
        }
    }
}
