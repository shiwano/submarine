using System;

namespace TyphenApi
{
    public static class Api
    {
        public static TyphenApi.WebApi.Submarine Submarine(string baseUri)
        {
            var uri = new Uri(baseUri);
            var controller = new Controller.Submarine();
            var sender = new WebApiRequestSenderWWW();
            var serializer = new JSONSerializer();
            return new TyphenApi.WebApi.Submarine(uri, controller, sender, serializer);
        }
    }
}
