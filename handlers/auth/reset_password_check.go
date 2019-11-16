package auth

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

type resetPasswordCheckBodyParam struct {
	ResetToken string `json:"token" validate:"required"`
}

// ResetPasswordCheck handles reset password check
func (auth Auth) ResetPasswordCheck(ctx *fasthttp.RequestCtx) {
	var param resetPasswordCheckBodyParam

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

	db := auth.appCtx.DB

	valid, errMsg := validateResetPasswordToken(db, param.ResetToken)

	if !valid {
		httpresponse.JSONErr(ctx, errMsg, fasthttp.StatusBadRequest)
		return
	}

	httpresponse.JSONOk(ctx, fasthttp.StatusOK)
}

func validateResetPasswordToken(db *gorm.DB, token string) (bool, string) {
	var resetCred models.ResetCredential
	if db.Where(&models.ResetCredential{
		ResetToken: token,
	}).First(&resetCred).RecordNotFound() {
		return false, "Invalid Token"
	}

	if resetCred.IsExpired {
		return false, "Token already expired"
	}

	now := time.Now()
	if !resetCred.ValidUntil.After(now) {
		db.Model(&resetCred).Updates(&models.ResetCredential{
			IsExpired: true,
		})

		return false, "Token already expired"
	}

	return true, ""
}
