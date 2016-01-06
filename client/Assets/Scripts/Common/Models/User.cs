using UnityEngine;
using System.Collections;
using Type = TyphenApi.Type.Submarine;
using Zenject;

namespace Submarine.Models
{
    public class User
    {
        public Type.LoggedInUser LoggedInUser { get; set; }
        public string ApiSessionKey { get; set; }

        public string Name { get { return LoggedInUser.Name; } }
        public bool IsInBattle { get { return LoggedInUser.JoinedRoom != null; } }
    }
}
