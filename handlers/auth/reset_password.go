package auth

import (
	"encoding/json"

	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

type resetPasswordBodyParam struct {
	ResetToken string `json:"token" validate:"required"`
	Email string `json:"email" validate:"required,email"`
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

	var cred models.Credential
	if db.Where(&models.Credential{
		Email: param.Email,
	}).First(&cred).RecordNotFound() {
		httpresponse.JSONErr(ctx, "User not found", fasthttp.StatusBadRequest)
		return
	}

	// get the email
	// check email on reset credentials if already exist
	// if not check email on credentials exist else send OK
	// send reset email
	httpresponse.JSON(ctx, nil, fasthttp.StatusOK)
}
