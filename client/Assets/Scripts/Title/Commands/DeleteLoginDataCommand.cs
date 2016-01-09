using UnityEngine;
using Zenject;
using Zenject.Commands;

namespace Submarine.Title
{
    public class DeleteLoginDataCommand : Command
    {
        public class Handler : ICommandHandler
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
