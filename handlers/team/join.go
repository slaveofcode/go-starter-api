package team

import (
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// JoinTeam will add invited member to the team
func (t Team) JoinTeam(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		// validate the token
		// check if user already joined to the team
		// add as member if not joined yet
	}
}
