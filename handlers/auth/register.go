package auth

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/password"
	"github.com/slaveofcode/go-starter-api/lib/random"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
)

type registerBodyParam struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CPassword string `json:"cpassword"`
}

type registerResponse struct {
}

// Register handles user registration
func (auth Auth) Register(ctx *fasthttp.RequestCtx) {
	// get data
	var param registerBodyParam
	err := json.Unmarshal(ctx.PostBody(), &param)
	if err != nil {
		httpresponse.JSONErr(ctx, "Wrong post data", fasthttp.StatusBadRequest)
		return
	}

	// check existing data
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
