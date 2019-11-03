package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// ForgotPassword handles user forgot password request
func (auth Auth) ForgotPassword(ctx *fasthttp.RequestCtx) {
	// check email exist on credential
	// if true send email 
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
