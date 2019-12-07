package subscription

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// CurrentPlan will return current running plan info
func (t Subscription) CurrentPlan(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// return current plan info
	}
}
