// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class Error : TyphenApi.TypeBase
    {
        [TyphenApi.SerializablePropertyAttribute("code", false)]
        public int Code { get; set; }
        [TyphenApi.SerializablePropertyAttribute("name", false)]
        public string Name { get; set; }
        [TyphenApi.SerializablePropertyAttribute("message", false)]
        public string Message { get; set; }
    }
}