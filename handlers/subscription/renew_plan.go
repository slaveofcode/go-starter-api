package subscription

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// Renew will renew existing plan
func (t Subscription) Renew(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// renew plan
	}
}
