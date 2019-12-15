package subscription

import (
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

// CurrentPlan will return current running plan info
func (t Subscription) CurrentPlan(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		db := t.appCtx.DB

		var plans []models.Subscription
		if err := db.Where(&models.Subscription{
			UserID: sessionData.UserID,
		}).Find(&plans).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to get plans: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		httpresponse.JSON(ctx, &plans, fasthttp.StatusOK)
		return
	}
}
