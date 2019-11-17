package middleware

import (
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	authSession "github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// SessionRequestHandler interface for handler with session
type SessionRequestHandler func(sessData *session.Data) func(ctx *fasthttp.RequestCtx)

// AuthenticatedUser provides authentication checking for user by identifying the session
func AuthenticatedUser(appCtx *context.AppContext, requestHandler SessionRequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		store, err := appCtx.Sesssion.Get(ctx)

		if err != nil {
			httpresponse.JSONErr(ctx, "Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		existingSess, err := authSession.GetAuth(store)

		if err != nil {
			httpresponse.JSONErr(ctx, "Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		if existingSess != nil {
			userSess := existingSess.(authSession.Data)
			if userSess.UserID != 0 {
				requestHandler(&userSess)(ctx)
				return
			}
		}

		httpresponse.JSONErr(ctx, "Unauthenticated", fasthttp.StatusForbidden)
		return
	}
}
