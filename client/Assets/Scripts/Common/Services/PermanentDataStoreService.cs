using UnityEngine;

namespace Submarine
{
    public class PermanentDataStoreService
    {
        public bool HasLoginData
        {
            get { return !string.IsNullOrEmpty(UserName) && !string.IsNullOrEmpty(Password); }
        }

        public string UserName
        {
            get { return PlayerPrefs.GetString("UserName"); }
            set { PlayerPrefs.SetString("UserName", value); }
        }

        public string Password
        {
            get { return PlayerPrefs.GetString("Password"); }
            set { PlayerPrefs.SetString("Password", value); }
        }

        public void Save()
        {
            PlayerPrefs.Save();
        }

        public void Clear()
        {
            PlayerPrefs.DeleteAll();
        }
    }
}
