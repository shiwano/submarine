using TyphenApi.Type.Submarine;

namespace TyphenApi.Controller.WebApi
{
    public class Submarine : IWebApiController<Error>
    {
        public IWebApiRequestSender RequestSender { get; private set; }
        public ISerializer RequestSerializer { get; private set; }
        public IDeserializer ResponseDeserializer { get; private set; }

        public Submarine()
        {
            RequestSender = new WebApiRequestSenderWWW();

            var jsonSerializer = new JSONSerializer();
            RequestSerializer = jsonSerializer;
            ResponseDeserializer = jsonSerializer;
        }

        public void OnBeforeRequestSend(IWebApiRequest request)
        {
            request.Headers["Content-Type"] = "application/json";
        }

        public void OnRequestError(WebApiError<Error> error)
        {
        }

        public void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response)
        {
        }
    }
}
