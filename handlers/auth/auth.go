package auth

import (
	"github.com/slaveofcode/go-starter-api/context"
)

// Auth class
type Auth struct {
	appCtx *context.AppContext
}

// NewAuth create new user instance
func NewAuth(appCtx *context.AppContext) *Auth {
	return &Auth{
		appCtx: appCtx,
	}
}
