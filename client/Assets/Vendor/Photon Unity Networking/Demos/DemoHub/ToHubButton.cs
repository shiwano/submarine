using UnityEngine;
using System.Collections;

public class ToHubButton : MonoBehaviour
{
    public Texture2D ButtonTexture;
    private Rect ButtonRect;

    private static ToHubButton instance;

    public static ToHubButton Instance
    {
        get
        {
            if (instance == null)
            {
                instance = FindObjectOfType(typeof(ToHubButton)) as ToHubButton;
            }

            return instance;
        }
    }

    void Awake()
    {
        if (Instance != null && Instance != this)
        {
            Destroy(this.gameObject);
        }
    }

	// Use this for initialization
	void Start () 
    {
        if (ButtonTexture == null)
        {
            this.gameObject.SetActive(false);
            return;
        }
	    DontDestroyOnLoad(this.gameObject);
	}

    public void OnGUI()
    {
        if (Application.loadedLevel != 0)
        {
            int w = ButtonTexture.width + 4;
            int h = ButtonTexture.height + 4;

            ButtonRect = new Rect(Screen.width - w, Screen.height - h, w, h);
            if (GUI.Button(ButtonRect, ButtonTexture, GUIStyle.none))
            {
                PhotonNetwork.Disconnect();
                Application.LoadLevel(0);
            }
        }
    }
}
