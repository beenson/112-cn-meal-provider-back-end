package handler

import (
	"net/http"

	"gitlab.winfra.cs.nycu.edu.tw/112-cn/meal-provider-back-end/pkg/gateway/internal"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Msg string `json:"msg"`
	}{"pong!"}

	internal.SendJSONResponse(w, resp)
}
