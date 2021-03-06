// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    [MessagePack.MessagePackObject]
    [Newtonsoft.Json.JsonObject(Newtonsoft.Json.MemberSerialization.OptIn)]
    public partial class Error : TyphenApi.TypeBase<Error>
    {
        [TyphenApi.QueryStringProperty("code", false)]
        [MessagePack.Key("code")]
        [Newtonsoft.Json.JsonProperty("code")]
        [Newtonsoft.Json.JsonRequired]
        public long Code { get; set; }
        [TyphenApi.QueryStringProperty("name", false)]
        [MessagePack.Key("name")]
        [Newtonsoft.Json.JsonProperty("name")]
        [Newtonsoft.Json.JsonRequired]
        public string Name { get; set; }
        [TyphenApi.QueryStringProperty("message", false)]
        [MessagePack.Key("message")]
        [Newtonsoft.Json.JsonProperty("message")]
        [Newtonsoft.Json.JsonRequired]
        public string Message { get; set; }
    }
}
