package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/valyala/fasthttp"
)

// Verify handles user verification
func (auth Auth) Verify(ctx *fasthttp.RequestCtx) {
	// get data
	// check the verification code exist
	// set VerifiedAt on user
	// register login session
	// redirect to dashboard
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
