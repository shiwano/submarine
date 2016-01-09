using UnityEngine;
using Zenject.Commands;

namespace Submarine.Title
{
    public class DeleteLoginDataCommand : Command
    {
        public class Handler : ICommandHandler
        {
            readonly PermanentDataStoreService dataStore;

            public Handler(PermanentDataStoreService dataStore)
            {
                this.dataStore = dataStore;
            }

            public void Execute()
            {
                dataStore.Clear();
                Debug.Log("Deleted login data");
            }
        }
    }
}
