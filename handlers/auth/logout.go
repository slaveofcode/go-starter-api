package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// Logout handles user logout action and clearing the login session
func (auth Auth) Logout(ctx *fasthttp.RequestCtx) {
	// check login sessino
	// remove the session
	// redirect to home
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
