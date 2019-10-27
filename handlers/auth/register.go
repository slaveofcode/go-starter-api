package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// Register handles user registration
func (auth Auth) Register(ctx *fasthttp.RequestCtx) {
	// get data
	// check existing data
	// create new user
	// create new credential
	// create user verification request
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
