using System;

namespace TyphenApi
{
    public static class Api
    {
        public static TyphenApi.WebApi.Submarine Submarine(string baseUri)
        {
            var uri = new Uri(baseUri);
            var controller = new TyphenApi.Controller.Submarine();
            return new TyphenApi.WebApi.Submarine(uri, controller);
        }
    }
}
