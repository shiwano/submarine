using UnityEngine;
using TyphenApi.Type.Submarine;
using Game = Submarine;
using System.Text.RegularExpressions;
using System;

namespace TyphenApi.WebApi
{
    public class Submarine : Base.Submarine
    {
        public class SessionKey
        {
            const string SessionKeyName = "_submarine_api_session";
            const string SessionKeyPattern = SessionKeyName + @"=([a-zA-Z0-9]+);";

            public readonly string Value;
            public string Cookie { get { return string.Format("{0}={1};", SessionKeyName, Value); } }

            public SessionKey(string setCookieValue)
            {
                var match = Regex.Match(setCookieValue, SessionKeyPattern, RegexOptions.Multiline);
                if (match == null)
                {
                    throw new InvalidOperationException("No session key!");
                }
                Value = match.Groups[1].Value;
            }
        }

        SessionKey sessionKey;

        public Submarine(string baseUri) : base(baseUri)
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
            if (sessionKey != null)
            {
                request.Headers["Cookie"] = sessionKey.Cookie;
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
                sessionKey = new SessionKey(setCookieValue);
                #if UNITY_EDITOR
                Game.Logger.LogWithColor(new Color(0.5f, 0f, 0.5f), "[WebAPI] SessionKey: " + sessionKey.Value, setCookieValue);
                #endif
            }
        }
    }
}
