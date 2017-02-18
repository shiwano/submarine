using UnityEngine;
using Zenject;

namespace Submarine.Title
{
    public class DeleteLoginDataCommand : Signal<DeleteLoginDataCommand>
    {
        public class Handler
        {
            [Inject]
            PermanentDataStoreService dataStore;

            public void Execute()
            {
                dataStore.Clear();
                Debug.Log("Deleted login data");
            }
        }
    }
}
