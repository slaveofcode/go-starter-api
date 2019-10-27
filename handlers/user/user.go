package user

import (
	"github.com/slaveofcode/go-starter-api/context"
)

// User class
type User struct {
	appCtx *context.AppContext
}

// NewUser create new user instance
func NewUser(appCtx *context.AppContext) *User {
	return &User{
		appCtx: appCtx,
	}
}
