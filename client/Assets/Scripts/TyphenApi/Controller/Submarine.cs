using TyphenApi.Type.Submarine;

namespace TyphenApi.Controller
{
    public class Submarine : IWebApiController<Error>
    {
        public IWebApiRequestSender RequestSender { get; private set; }
        public ISerializer Serializer { get; private set; }

        public Submarine()
        {
            RequestSender = new WebApiRequestSenderWWW();
            Serializer = new JSONSerializer();
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
