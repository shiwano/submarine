using TyphenApi.Type.Submarine;

namespace TyphenApi.Controller
{
    public class Submarine : IWebApiController<Error>
    {
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
