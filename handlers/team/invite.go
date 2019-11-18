package team

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// InviteMember send invitation to join to the team
func (t Team) InviteMember(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// get email of member
		// get role
		// create token
		// send email with token link
	}
}
