package auth

import (
	"encoding/json"
	"time"

	"github.com/fasthttp/session"
	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/password"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

type loginBodyParam struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// SessionData wrapper data for login session
type SessionData struct {
	UserID    uint
	Name      string
	Email     string
	LoginTime time.Time
}

// Login handle login user and creates the session
func (auth Auth) Login(ctx *fasthttp.RequestCtx) {
	db := auth.appCtx.DB
	store, err := auth.appCtx.Sesssion.Get(ctx)

	if err != nil {
		ctx.Error("Internal server error: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	defer auth.appCtx.Sesssion.Save(ctx, store)

	existingSess := getSessionAuth(store)
	if existingSess == nil {
		logrus.Info("session nil")
	} else {
		logrus.Info("session not nil")
	}

	if existingSess != nil {
		userSess := existingSess.(SessionData)
		if userSess.UserID != 0 {
			httpresponse.JSONOk(ctx, fasthttp.StatusOK)
			return
		}
	}

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
		httpresponse.JSONErr(ctx, "User not verified", fasthttp.StatusBadRequest)
		return
	}

	setSessionAuth(store, &SessionData{
		UserID:    cred.UserID,
		Name:      user.Name,
		Email:     cred.Email,
		LoginTime: time.Now(),
	})

	httpresponse.JSONOk(ctx, fasthttp.StatusOK)
	return
}

func setSessionAuth(store session.Storer, data *SessionData) {
	store.Set("auth_id", data.UserID)
	store.Set("auth_name", data.Name)
	store.Set("auth_email", data.Email)
	store.Set("auth_login_time", data.LoginTime)
}
func getSessionAuth(store session.Storer) interface{} {
	authData := store.Get("auth")

	if authData != nil {
		data := authData.(SessionData)
		return SessionData{
			UserID:    data.UserID,
			Name:      data.Name,
			Email:     data.Email,
			LoginTime: data.LoginTime,
		}
	}

	return authData
}
