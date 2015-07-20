using System.Collections;

namespace TyphenApi
{
    public interface IWebApiRequestSender
    {
        IEnumerator Send(IWebApiRequest request);
    }
}
