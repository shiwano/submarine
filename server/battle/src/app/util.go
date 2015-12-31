package main

import (
	"net/http"
	"strconv"
)

func getSessionID(request *http.Request) (uint64, error) {
	return strconv.ParseUint(request.URL.Query().Get("session_id"), 10, 64)
}

func getBattleID(request *http.Request) (uint64, error) {
	return strconv.ParseUint(request.URL.Query().Get("battle_id"), 10, 64)
}
