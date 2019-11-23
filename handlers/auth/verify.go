package auth

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	emailMedia = "EMAIL"
)

type verifyBodyParam struct {
	Media string `json:"media" validate:"required"`
	Token string `json:"token" validate:"required"`
}

// Verify handles user verification
func (auth Auth) Verify(ctx *fasthttp.RequestCtx) {
	db := auth.appCtx.DB
	var param verifyBodyParam
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

	if !isValidMedia(param.Media) {
		httpresponse.JSONErr(ctx, "Invalid media: "+param.Media+":"+param.Token, fasthttp.StatusBadRequest)
		return
	}

	var userVerify models.UserVerificationRequest
	if !db.Where(&models.UserVerificationRequest{
		Type:            param.Media,
		VerificationKey: param.Token,
	}).First(&userVerify).RecordNotFound() {
		tx := db.Begin()
		if userVerify.VerifiedAt != nil {
			httpresponse.JSONOk(ctx, fasthttp.StatusOK)
			return
		}

		t := time.Now()
		err = tx.Model(&userVerify).Updates(models.UserVerificationRequest{
			VerifiedAt: &t,
		}).Error

		if err != nil {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "Unable to verify token: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		var user models.User
		if tx.Where(&models.User{
			ID: userVerify.UserID,
		}).First(&user).RecordNotFound() {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "User not found: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		if err = tx.Model(&user).Updates(models.User{
			VerifiedAt: &t,
		}).Error; err != nil {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "Unable to verify token: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		if err := tx.Create(&models.UserVerificationAttempt{
			UserID:                    userVerify.UserID,
			UserVerificationRequestID: userVerify.ID,
			UserAgent:                 string(ctx.UserAgent()),
			IPAddr:                    ctx.RemoteAddr().String(),
		}).Error; err != nil {
			tx.Rollback()
			logrus.Error("Unable to record verification attempt")
		}

		tx.Commit()

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}

	httpresponse.JSONErr(ctx, "Token not found", fasthttp.StatusNotFound)
	return
}

func isValidMedia(media string) bool {
	if media == emailMedia {
		return true
	}

	return false
}
