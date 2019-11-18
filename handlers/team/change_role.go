package team

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// ChangeMemberRole will change role of member team
func (t Team) ChangeMemberRole(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// get current user role, check if allowed to make a change
		// get the member
		// get the role
		// do change
	}
}
