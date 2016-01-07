using UnityEngine;
using Zenject.Commands;

namespace Submarine.Commands
{
    public class DeleteLoginData : Command
    {
        public class Handler : ICommandHandler
        {
            readonly Services.PermanentDataStore dataStore;

            public Handler(Services.PermanentDataStore dataStore)
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
