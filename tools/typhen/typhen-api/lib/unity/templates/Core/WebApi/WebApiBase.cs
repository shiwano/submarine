using System;

namespace TyphenApi
{
    public abstract class WebApiBase<ErrorT> where ErrorT : TypeBase
    {
        public Uri BaseUri { get; private set; }
        public IWebApiController<ErrorT> Controller { get; protected set; }
        public IWebApiRequestSender RequestSender { get { return Controller.RequestSender; } }
        public ISerializer Serializer { get { return Controller.Serializer; } }

        protected WebApiBase(Uri baseUri, IWebApiController<ErrorT> controller)
        {
            BaseUri = baseUri;
            Controller = controller;
        }

        public WebApiError<ErrorT> Error(Exception error)
        {
            return error as WebApiError<ErrorT>;
        }
    }
}
