package main

import (
	"app/typhenapi/core"
	webapi "app/typhenapi/web/submarine"
	"net/http"
)

// NewWebAPI creates a submarine WebAPI instance.
func NewWebAPI(baseURI string) *webapi.WebAPI {
	serializer := typhenapi.NewJSONSerializer()
	api := webapi.New(baseURI, serializer)
	api.OnBeforeRequestSend = onBeforeWebAPIRequestSend
	return api
}

func onBeforeWebAPIRequestSend(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
}
