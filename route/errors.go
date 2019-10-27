package route

import (
	log "github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// ErrorResponse used for JSON error response
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

// NotFoundHandler will used as 404 handler
func NotFoundHandler(ctx *fasthttp.RequestCtx) {
	httpresponse.JSON(ctx, &ErrorResponse{
		Error:   "NotFound",
		Message: "Requested resource are not found",
	}, fasthttp.StatusNotFound)
}

// PanicHandler will used as panic handler
func PanicHandler(ctx *fasthttp.RequestCtx, err interface{}) {
	log.Error("Panic Error: ", err)
	log.Error("URI: ", ctx.Request.URI())
	httpresponse.JSON(ctx, &ErrorResponse{
		Error:   "FatalError",
		Message: "Fatal error reached",
	}, fasthttp.StatusInternalServerError)
}
