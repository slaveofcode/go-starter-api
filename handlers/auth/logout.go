package auth

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/valyala/fasthttp"
)

// Logout handles user logout action and clearing the login session
func (auth Auth) Logout(ctx *fasthttp.RequestCtx) {
	store, err := auth.appCtx.Session.Get(ctx)

	if err != nil {
		ctx.Error("Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer auth.appCtx.Session.Save(ctx, store)

	existingSess, err := session.GetAuth(store)

	if err != nil {
		ctx.Error("Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if existingSess != nil {
		store.Delete("auth")
		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}

	httpresponse.JSONFail(ctx, fasthttp.StatusBadRequest)
	return
}
