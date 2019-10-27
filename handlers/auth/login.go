package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// Login handle login user and creates the session
func (auth Auth) Login(ctx *fasthttp.RequestCtx) {
	// get email & password
	// check email exist on credential
	// verify the password with the given one
	// register session user
	// redirect dashboard / to referrer url if exist
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
