using System;

namespace TyphenApi
{
    public abstract class WebApiBase<ErrorT> where ErrorT : TypeBase
    {
        public Uri BaseUri { get; private set; }
        public IWebApiController<ErrorT> Controller { get; private set; }
        public IWebApiRequestSender RequestSender { get; private set; }
        public ISerializer Serializer { get; private set; }

        protected WebApiBase(Uri baseUri, IWebApiController<ErrorT> controller,
            IWebApiRequestSender requestSender, ISerializer serializer)
        {
            BaseUri = baseUri;
            Controller = controller;
            RequestSender = requestSender;
            Serializer = serializer;
        }

        public WebApiError<ErrorT> Error(Exception error)
        {
            return error as WebApiError<ErrorT>;
        }
    }
}
