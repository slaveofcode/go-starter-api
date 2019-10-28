package auth

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/password"
	"github.com/slaveofcode/go-starter-api/lib/random"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"

	validator "gopkg.in/go-playground/validator.v9"
)

type registerBodyParam struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	CPassword string `json:"cpassword" validate:"required"`
}

type registerResponse struct {
}

// Register handles user registration
func (auth Auth) Register(ctx *fasthttp.RequestCtx) {
	var param registerBodyParam
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

	if param.Password != param.CPassword {
		httpresponse.JSONErr(ctx, "Password doesn't match", fasthttp.StatusBadRequest)
		return
	}

	var credential models.Credential
	if auth.appCtx.DB.Where(&models.Credential{
		Email: param.Email,
	}).First(&credential).RecordNotFound() {
		tx := auth.appCtx.DB.Begin()

		user, err := createUser(tx, &param)
		if err != nil {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "Unable to process registration: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = createCredential(tx, user, &param)
		if err != nil {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "Unable to process registration: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		err = createUserVerificationRequest(tx, user, &param)
		if err != nil {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "Unable to process registration: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		tx.Commit()
		httpresponse.JSONOk(ctx, fasthttp.StatusCreated)
		return
	}

	httpresponse.JSON(ctx, &credential, fasthttp.StatusCreated)
	return
}

func createUser(db *gorm.DB, p *registerBodyParam) (*models.User, error) {
	user := models.User{
		Name:           p.Name,
		Timezone:       "Asia/Jakarta",
		TimezoneOffset: "+7",
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
func createCredential(db *gorm.DB, user *models.User, p *registerBodyParam) error {
	hashed, err := password.Hash(p.Password)

	if err != nil {
		return err
	}

	return db.Create(&models.Credential{
		Email:    p.Email,
		Password: hashed,
		UserID:   user.ID,
	}).Error
}
func createUserVerificationRequest(db *gorm.DB, user *models.User, p *registerBodyParam) error {
	return db.Create(&models.UserVerificationRequest{
		UserID:          user.ID,
		Type:            "EMAIL",
		VerificationKey: random.GetStr(100),
	}).Error
}
