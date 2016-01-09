using UnityEngine;
using TyphenApi.Type.Submarine;
using Game = Submarine;

namespace TyphenApi.WebApi
{
    public class Submarine : Base.Submarine
    {
        public Submarine(string baseUri) : base(baseUri)
        {
            RequestSender = new WebApiRequestSenderWWW();

            var jsonSerializer = new JSONSerializer();
            RequestSerializer = jsonSerializer;
            ResponseDeserializer = jsonSerializer;
        }

        public override void OnBeforeRequestSend(IWebApiRequest request)
        {
            request.Headers["Content-Type"] = "application/json";
            #if UNITY_EDITOR
            Game.Logger.LogWithColor(Color.blue, "[WebAPI] Request: " + request.Uri, request.Body);
            #endif
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
        }
    }
}
