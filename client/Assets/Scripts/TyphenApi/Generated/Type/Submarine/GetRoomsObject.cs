// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    [MessagePack.MessagePackObject]
    [Newtonsoft.Json.JsonObject(Newtonsoft.Json.MemberSerialization.OptIn)]
    public partial class GetRoomsObject : TyphenApi.TypeBase<GetRoomsObject>
    {
        [TyphenApi.QueryStringProperty("rooms", false)]
        [MessagePack.Key("rooms")]
        [Newtonsoft.Json.JsonProperty("rooms")]
        [Newtonsoft.Json.JsonRequired]
        public List<TyphenApi.Type.Submarine.Room> Rooms { get; set; }
    }
}
