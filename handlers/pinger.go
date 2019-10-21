package handlers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// Pinger send ping result as server running indication
func Pinger(ctx *fasthttp.RequestCtx) {
	panic("Error Here from Panic")
	res := struct {
		Status string `json:"status"`
		Live   bool   `json:"isAlive"`
	}{
		Status: "OK",
		Live:   true,
	}
	resJSON, _ := json.Marshal(res)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(resJSON)
}
