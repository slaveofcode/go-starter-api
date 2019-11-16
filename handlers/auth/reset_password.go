package auth

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/password"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

type resetPasswordBodyParam struct {
	ResetToken string `json:"token" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

// ResetPassword handles reset password action and change the old credential with the new one
func (auth Auth) ResetPassword(ctx *fasthttp.RequestCtx) {
	var param resetPasswordBodyParam

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

	var resetCred models.ResetCredential
	if db.Preload("Credential").Where(&models.ResetCredential{
		ResetToken: param.ResetToken,
	}).First(&resetCred).RecordNotFound() {
		httpresponse.JSONErr(ctx, "Token not found", fasthttp.StatusBadRequest)
		return
	}

	if resetCred.ValidatedAt != nil {
		httpresponse.JSONErr(ctx, "Token not found", fasthttp.StatusBadRequest)
		return
	}

	tx := auth.appCtx.DB.Begin()

	var user models.User
	if tx.Where(&models.User{
		ID: resetCred.Credential.UserID,
	}).First(&user).RecordNotFound() {
		tx.Rollback()
		httpresponse.JSONErr(ctx, "Unable to reset password", fasthttp.StatusInternalServerError)
		return
	}

	if err = resetCredentialPassword(tx, &resetCred.Credential, param.Password); err != nil {
		tx.Rollback()
		httpresponse.JSONErr(ctx, "Unable to reset password", fasthttp.StatusInternalServerError)
		return
	}

	now := time.Now()
	tx.Model(&resetCred).Updates(&models.ResetCredential{
		ValidatedAt: &now,
	})

	tx.Commit()

	httpresponse.JSONOk(ctx, fasthttp.StatusOK)
}

func resetCredentialPassword(db *gorm.DB, credential *models.Credential, p string) error {
	hashed, err := password.Hash(p)

	if err != nil {
		return err
	}

	return db.Model(credential).Updates(&models.Credential{
		Password: hashed,
	}).Error
}
