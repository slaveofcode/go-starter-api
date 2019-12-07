package subscription

import "github.com/slaveofcode/go-starter-api/context"

// Subscription handler
type Subscription struct {
	appCtx *context.AppContext
}

// NewSubscription create new subscription instance
func NewSubscription(appCtx *context.AppContext) *Subscription {
	return &Subscription{
		appCtx: appCtx,
	}
}
