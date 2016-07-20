using UnityEngine;

namespace Submarine
{
    public class PermanentDataStoreService
    {
        public bool HasSignedUp
        {
            get { return !string.IsNullOrEmpty(AuthToken); }
        }

        public string AuthToken
        {
            get { return PlayerPrefs.GetString("AuthToken"); }
            set { PlayerPrefs.SetString("AuthToken", value); }
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
