package subscription

import (
	"encoding/json"
	"time"

	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type startPlanBodyParam struct {
	Plan           string `json:"plan" validate:"required"`
	DurationInDays int    `json:"durationInDays" validate:"required"`
	IsRecurring    bool   `json:"isRecurring"`
}

// StartPlan will create a new plan
func (t Subscription) StartPlan(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param startPlanBodyParam
		err := json.Unmarshal(ctx.PostBody(), &param)
		if err != nil {
			httpresponse.JSONErr(ctx, "Wrong post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		v := validator.New()
		err = v.Struct(param)
		if err != nil {
			httpresponse.JSONErr(ctx, "Invalid post data: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if IsValidPlan(param.Plan) {
			httpresponse.JSONErr(ctx, "Invalid Plan: "+param.Plan, fasthttp.StatusBadRequest)
			return
		}

		db := t.appCtx.DB

		var currSub models.Subscription
		if err := db.Where(&models.Subscription{
			UserID:   sessionData.UserID,
			PlanType: param.Plan,
		}).First(&currSub).Error; err != nil {
			httpresponse.JSONErr(ctx, "Unable to check existing plan: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if currSub.ID != 0 && currSub.EndOfPlanAt.Before(time.Now()) {
			// create new plan
			if err = db.Create(&models.Subscription{
				UserID:      sessionData.UserID,
				PlanType:    param.Plan,
				EndOfPlanAt: time.Now().AddDate(0, 0, param.DurationInDays),
				IsRecurring: param.IsRecurring,
			}).Error; err != nil {
				httpresponse.JSONErr(ctx, "Unable to create plan: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}
		} else {
			// extend the plan
			if err = db.Model(&currSub).Updates(&models.Subscription{
				EndOfPlanAt: currSub.EndOfPlanAt.AddDate(0, 0, param.DurationInDays),
				IsRecurring: param.IsRecurring,
			}).Error; err != nil {
				httpresponse.JSONErr(ctx, "Unable to extend plan: "+err.Error(), fasthttp.StatusInternalServerError)
				return
			}
		}

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
	}
}
