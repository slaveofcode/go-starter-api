package subscription

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// StopPlan will stop existing plan for renewal
func (t Subscription) StopPlan(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// stop plan
	}
}
