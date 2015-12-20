using TyphenApi.Type.Submarine;

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
        }

        public override void OnRequestError(IWebApiRequest request, WebApiError<Error> error)
        {
        }

        public override void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response)
        {
        }
    }
}
