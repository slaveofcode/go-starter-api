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

// IsValidPlan will validate plan
func IsValidPlan(plan string) bool {
	if plan == PlanFree {
		return true
	}

	if plan == PlanStarter {
		return true
	}

	if plan == PlanPremium {
		return true
	}

	return false
}

func PlanToHuman(plan string) string {
	if plan == PlanFree {
		return "Free Plan"
	}

	if plan == PlanStarter {
		return "Starter Plan"
	}

	if plan == PlanPremium {
		return "Premium Plan"
	}

	return ""
}

func PlanDescriptionToHuman(plan string) string {
	if plan == PlanFree {
		return "Go Starter - Free Plan Subscription"
	}

	if plan == PlanStarter {
		return "Go Starter - Starter Plan Subscription"
	}

	if plan == PlanPremium {
		return "Go Starter - Premium Plan Subscription"
	}

	return ""
}

func PlanToPriceUSD(plan string) int64 {
	if plan == PlanFree {
		return 0
	}

	if plan == PlanStarter {
		return 50
	}

	if plan == PlanPremium {
		return 100
	}

	return 100
}

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
