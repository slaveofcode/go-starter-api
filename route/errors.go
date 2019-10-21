package route

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

// ErrorResponse used for JSON error response
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

// NotFoundHandler will used as 404 handler
func NotFoundHandler(ctx *fasthttp.RequestCtx) {
	res := ErrorResponse{
		Error:   "NotFound",
		Message: "Requested resource are not found",
	}
	resJSON, _ := json.Marshal(res)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBody(resJSON)
}

// PanicHandler will used as panic handler
func PanicHandler(ctx *fasthttp.RequestCtx, err interface{}) {
	log.Error("Panic Error: ", err)
	log.Error("URI: ", ctx.Request.URI())
	res := ErrorResponse{
		Error:   "FatalError",
		Message: "Fatal error reached",
	}
	resJSON, _ := json.Marshal(res)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	ctx.SetBody(resJSON)
}
