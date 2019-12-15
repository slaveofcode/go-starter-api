package subscription

import "github.com/slaveofcode/go-starter-api/context"

const (
	// PlanFree free plan
	PlanFree = "PLAN_FREE"
	// PlanStarter starter plan
	PlanStarter = "PLAN_STARTER"
	// PlanPremium premium plan
	PlanPremium = "PLAN_PREMIUM"
)

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
