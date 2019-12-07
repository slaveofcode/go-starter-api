package subscription

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// StartPlan will create a new plan
func (t Subscription) StartPlan(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// create a new plan
	}
}
