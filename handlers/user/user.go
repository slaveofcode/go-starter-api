package user

import (
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
)

// User class
type User struct {
	appCtx *context.AppContext
}

type listResponse struct {
	Items []models.User `json:"items"`
	Total int           `json:"total"`
}

// NewUser create new user instance
func NewUser(appCtx *context.AppContext) *User {
	return &User{
		appCtx: appCtx,
	}
}
