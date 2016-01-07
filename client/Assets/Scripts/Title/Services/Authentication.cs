using System;
using System.Text.RegularExpressions;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Services
{
    public class Authentication
    {
        public struct Result
        {
            public readonly string SessionKey;
            public readonly Type.LoggedInUser LoggedInUser;

            public Result(string sessionKey, Type.LoggedInUser loggedInUser)
            {
                SessionKey = sessionKey;
                LoggedInUser = loggedInUser;
            }
        }

        const string SessionKeyPattern = @"_submarine_api_session=([a-zA-Z0-9]+);";

        readonly TyphenApi.WebApi.Submarine webApi;

        public Authentication(TyphenApi.WebApi.Submarine webApi)
        {
            this.webApi = webApi;
        }

        public IObservable<Result> Login(string userName, string passWord)
        {
            return webApi.Login(userName, passWord)
                .Send()
                .Select(r =>
                {
                    var sessionKey = GetSessionKeyFromCookie(r.Headers["Set-Cookie"]);
                    return new Result(sessionKey, r.Data.User);
                });
        }

        public IObservable<Result> SignUp(string userName, string passWord)
        {
            return webApi.SignUp(userName, passWord)
                .Send()
                .Select(r =>
                {
                    var sessionKey = GetSessionKeyFromCookie(r.Headers["SET-COOKIE"]);
                    return new Result(sessionKey, r.Data.User);
                });
        }

        string GetSessionKeyFromCookie(string cookie)
        {
            var match = Regex.Match(cookie, SessionKeyPattern, RegexOptions.Multiline);
            if (match == null)
            {
                throw new InvalidOperationException("No session key!");
            }
            return match.Groups[1].Value;
        }
    }
}
