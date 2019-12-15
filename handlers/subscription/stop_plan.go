package subscription

import (
	"encoding/json"

	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type endPlanBodyParam struct {
	Plan string `json:"plan" validate:"required"`
}

// StopPlan will stop existing plan for renewal
func (t Subscription) StopPlan(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		var param endPlanBodyParam
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
			httpresponse.JSONErr(ctx, "Unable to get existing plan: "+err.Error(), fasthttp.StatusBadRequest)
			return
		}

		if currSub.ID == 0 {
			httpresponse.JSONErr(ctx, "No Plan available", fasthttp.StatusBadRequest)
			return
		}

		if currSub.IsRecurring {
			if err = db.Model(&currSub).Update("isRecurring", false).Error; err != nil {
				httpresponse.JSONErr(ctx, "Unable to stop existing plan: "+err.Error(), fasthttp.StatusBadRequest)
				return
			}
		}

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}
}
