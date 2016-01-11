using UnityEngine;
using TyphenApi.Type.Submarine;
using Game = Submarine;
using System.Text.RegularExpressions;

namespace TyphenApi.WebApi
{
    public class Submarine : Base.Submarine
    {
        class SessionId
        {
            const string SessionIdName = "_submarine_api_session";
            const string SessionIdPattern = SessionIdName + @"=([a-zA-Z0-9]+);";
            public string Id { get; private set; }
            public bool IsExists { get { return !string.IsNullOrEmpty(Id); } }

            public string Cookie
            {
                get
                {
                    return IsExists ?
                        string.Format("{0}={1};", SessionIdName, Id) :
                        string.Empty;
                }
            }

            public bool Update(string setCookieValue)
            {
                var oldId = Id;
                var match = Regex.Match(setCookieValue, SessionIdPattern, RegexOptions.Multiline);
                if (match != null)
                {
                    Id = match.Groups[1].Value;
                }
                return oldId != Id;
            }
        }

        readonly SessionId sessionId = new SessionId();

        public Submarine(Config config) : base(config.WebApiServerBaseUri)
        {
            RequestSender = new WebApiRequestSenderWWW();

            var jsonSerializer = new JSONSerializer();
            RequestSerializer = jsonSerializer;
            ResponseDeserializer = jsonSerializer;
        }

        public override void OnBeforeRequestSend(IWebApiRequest request)
        {
            #if UNITY_EDITOR
            Game.Logger.LogWithColor(Color.blue, "[WebAPI] Request: " + request.Uri, request.Body);
            #endif

            request.Headers["Content-Type"] = "application/json";
            if (sessionId.IsExists)
            {
                request.Headers["X-Cookie"] = sessionId.Cookie;
            }
        }

        public override void OnRequestError(IWebApiRequest request, WebApiError<Error> error)
        {
            #if UNITY_EDITOR
            Game.Logger.LogError("[WebAPI] Error: " + request.Uri, error.RawErrorMessage, error.Error);
            #endif
        }

        public override void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response)
        {
            #if UNITY_EDITOR
            Game.Logger.LogWithColor(new Color(0f, 0.4f, 0f), "[WebAPI] Response: " + request.Uri, response.Body);
            #endif

            string setCookieValue;
            if (response.Headers.TryGetValue("SET-COOKIE", out setCookieValue))
            {
                if (sessionId.Update(setCookieValue))
                {
                    #if UNITY_EDITOR
                    Game.Logger.LogWithColor(new Color(0.5f, 0f, 0.5f), "[WebAPI] SessionKey: " + sessionId.Id, setCookieValue);
                    #endif
                }
            }
        }
    }
}
