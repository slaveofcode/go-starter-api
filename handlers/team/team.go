package team

import (
	"github.com/slaveofcode/go-starter-api/context"
)

const (
	Owner   string = "Owner"
	Manager string = "Manager"
	Staff   string = "Staff"
)

// Team class
type Team struct {
	appCtx *context.AppContext
}

// NewTeam create new team instance
func NewTeam(appCtx *context.AppContext) *Team {
	return &Team{
		appCtx: appCtx,
	}
}
