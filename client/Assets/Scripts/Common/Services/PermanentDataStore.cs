using UnityEngine;
using System.Collections;

namespace Submarine.Services
{
    public class PermanentDataStore
    {
        public bool HasSigned
        {
            get { return string.IsNullOrEmpty(UserName) || string.IsNullOrEmpty(Password); }
        }

        public string ApiSessionKey
        {
            get { return PlayerPrefs.GetString("ApiSessionKey"); }
            set { PlayerPrefs.SetString("ApiSessionKey", value); }
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
