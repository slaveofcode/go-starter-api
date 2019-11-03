package httpresponse

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type jsonErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"msg"`
}

type jsonStatusResponse struct {
	Status bool `json:"status"`
}

func makeErr(msg string) []byte {
	resErr, _ := json.Marshal(&jsonErrorResponse{
		Error:   true,
		Message: msg,
	})
	return resErr
}

func makeStatus(s bool) []byte {
	resStat, _ := json.Marshal(&jsonStatusResponse{
		Status: s,
	})
	return resStat
}

// JSON make json response
func JSON(ctx *fasthttp.RequestCtx, body interface{}, status int) {
	ctx.SetContentType("application/json")

	out, err := json.Marshal(body)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(makeErr("Unable to process response: " + err.Error()))
	} else {
		ctx.SetStatusCode(status)
		ctx.SetBody(out)
	}
}

// JSONOk return positive response status
func JSONOk(ctx *fasthttp.RequestCtx, status int) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)
	ctx.SetBody(makeStatus(true))
}

// JSONFail return negative response status
func JSONFail(ctx *fasthttp.RequestCtx, status int) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)
	ctx.SetBody(makeStatus(false))
}

// JSONErr make error json response
func JSONErr(ctx *fasthttp.RequestCtx, message string, status int) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)
	ctx.SetBody(makeErr(message))
}
