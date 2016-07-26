using UnityEngine;
using TyphenApi.Type.Submarine;
using Game = Submarine;

namespace TyphenApi.WebApi
{
    public class Submarine : Base.Submarine
    {
        string accessToken;

        public bool IsAuthenticated { get { return !string.IsNullOrEmpty(accessToken); } }

        public Submarine(Config config) : base(config.ApiServerBaseUri)
        {
            RequestSender = new WebApiRequestSenderUnity();

            var jsonSerializer = new JSONSerializer();
            RequestSerializer = jsonSerializer;
            ResponseDeserializer = jsonSerializer;
        }

        public override void OnBeforeRequestSend(IWebApiRequest request)
        {
            #if UNITY_EDITOR
            Game.Logger.LogWithBlue("[WebAPI] Request: " + request.Uri, request.Body);
            #endif

            request.Headers["Content-Type"] = "application/json";

            if (IsAuthenticated)
            {
                request.Headers["X-Access-Token"] = accessToken;
            }
        }

        public override void OnRequestError(IWebApiRequest request, WebApiError<Error> error)
        {
            #if UNITY_EDITOR
            if (error.Error != null)
            {
                Game.Logger.LogError("[WebAPI] Error: " + request.Uri, error.Error.Code + ": " + error.Error.Name, error.Error);
            }
            else
            {
                Game.Logger.LogError("[WebAPI] Error: " + request.Uri, error.RawErrorMessage);
            }
            #endif
        }

        public override void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response)
        {
            #if UNITY_EDITOR
            Game.Logger.LogWithGreen("[WebAPI] Response: " + request.Uri, response.Body);
            #endif

            string accessToken;
            if (response.Headers.TryGetValue("X-Set-Access-Token", out accessToken))
            {
                this.accessToken = accessToken;

                #if UNITY_EDITOR
                Game.Logger.LogWithPurple("[WebAPI] AccessToken: " + accessToken);
                #endif
            }
        }
    }
}
