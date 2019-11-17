package user

import (
	"github.com/slaveofcode/go-starter-api/lib/hashids"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

const referralMinDigit = 5

type makeReferralResponseBody struct {
	Code string `json:"referralCode"`
}

// MakeReferral will create new referral code
func (u User) MakeReferral(sessionData *session.Data) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		db := u.appCtx.DB

		var referralCode models.ReferralCode
		if db.Where(&models.ReferralCode{
			UserID: sessionData.UserID,
		}).First(&referralCode).RecordNotFound() {
			h := hashids.New(referralMinDigit)

			code, err := h.Encode([]int{int(sessionData.UserID)})
			if err != nil {
				httpresponse.JSONErr(ctx, "Unable to generate referral code", fasthttp.StatusInternalServerError)
				return
			}

			if err := db.Create(&models.ReferralCode{
				UserID: sessionData.UserID,
				Code:   code,
			}).Error; err != nil {
				httpresponse.JSONErr(ctx, "Unable to generate referral code", fasthttp.StatusInternalServerError)
				return
			}

			httpresponse.JSON(ctx, makeReferralResponseBody{
				Code: code,
			}, fasthttp.StatusOK)
			return
		}

		httpresponse.JSON(ctx, makeReferralResponseBody{
			Code: referralCode.Code,
		}, fasthttp.StatusOK)
		return
	}
}
