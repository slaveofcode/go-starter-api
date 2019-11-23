package auth

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/password"
	authSession "github.com/slaveofcode/go-starter-api/lib/session"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

type loginBodyParam struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login handle login user and creates the session
func (auth Auth) Login(ctx *fasthttp.RequestCtx) {
	db := auth.appCtx.DB
	store, err := auth.appCtx.Session.Get(ctx)

	if err != nil {
		ctx.Error("Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer auth.appCtx.Session.Save(ctx, store)

	var param loginBodyParam

	if err := json.Unmarshal(ctx.PostBody(), &param); err != nil {
		httpresponse.JSONErr(ctx, "Wrong post data: "+err.Error(), fasthttp.StatusBadRequest)
		return
	}

	v := validator.New()
	if err := v.Struct(param); err != nil {
		httpresponse.JSONErr(ctx, "Invalid post data: "+err.Error(), fasthttp.StatusBadRequest)
		return
	}

	existingSess, err := authSession.GetAuth(store)

	if err != nil {
		ctx.Error("Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if existingSess == nil {
		logrus.Info("No session defined")
	} else {
		logrus.Info("Existing session data was found!")
	}

	if existingSess != nil {
		userSess := existingSess.(authSession.Data)
		if userSess.UserID != 0 && userSess.Email == param.Email {
			httpresponse.JSONOk(ctx, fasthttp.StatusOK)
			return
		}
	}

	var cred models.Credential
	if db.Where(&models.Credential{
		Email: param.Email,
	}).First(&cred).RecordNotFound() {
		httpresponse.JSONErr(ctx, "Email or Password not match", fasthttp.StatusBadRequest)
		return
	}

	if !password.Compare(param.Password, cred.Password) {
		httpresponse.JSONErr(ctx, "Email or Password not match", fasthttp.StatusBadRequest)
		return
	}

	var user models.User
	if db.Where(&models.User{
		ID: cred.UserID,
	}).First(&user).RecordNotFound() {
		httpresponse.JSONErr(ctx, "User not found", fasthttp.StatusBadRequest)
		return
	}

	if user.VerifiedAt == nil {
		httpresponse.JSONErr(ctx, "User not verified", fasthttp.StatusBadRequest)
		return
	}

	err = authSession.SetAuth(store, &authSession.Data{
		UserID:    cred.UserID,
		Name:      user.Name,
		Email:     cred.Email,
		LoginTime: time.Now(),
	})

	if err != nil {
		ctx.Error("Unable to make session: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	httpresponse.JSONOk(ctx, fasthttp.StatusOK)
	return
}
