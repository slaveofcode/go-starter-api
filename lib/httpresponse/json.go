package httpresponse

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// JSON make json response
func JSON(ctx *fasthttp.RequestCtx, body interface{}, status int) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(status)

	resUser, _ := json.Marshal(body)
	ctx.SetBody(resUser)
}
