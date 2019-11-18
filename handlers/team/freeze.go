package team

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// FreezeMember will freeze the member which disabling all actions
func (t Team) FreezeMember(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// get current user role, check if allowed to make a freeze
		// get the member
		// do freeze
	}
}
