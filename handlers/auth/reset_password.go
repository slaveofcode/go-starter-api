package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// ResetPassword handles reset password action and change the old credential with the new one
func (auth Auth) ResetPassword(ctx *fasthttp.RequestCtx) {
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
