package auth

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/matcornic/hermes"
	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/mail"
	"github.com/slaveofcode/go-starter-api/lib/password"
	"github.com/slaveofcode/go-starter-api/lib/random"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"

	validator "gopkg.in/go-playground/validator.v9"
)

type registerBodyParam struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	CPassword    string `json:"cpassword" validate:"required"`
	ReferralCode string `json:"referralCode"`
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

		verifyKey, err := createUserVerificationRequest(tx, user, &param)
		if err != nil {
			tx.Rollback()
			httpresponse.JSONErr(ctx, "Unable to process registration: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		if param.ReferralCode != "" {
			// Ignores error if any
			defineReferralUser(tx, param.ReferralCode, user)
		}

		tx.Commit()

		sendEmailVerification(param.Name, param.Email, verifyKey)

		httpresponse.JSONOk(ctx, fasthttp.StatusCreated)
		return
	}

	httpresponse.JSONErr(ctx, "Already registered", fasthttp.StatusBadRequest)
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
func createUserVerificationRequest(db *gorm.DB, user *models.User, p *registerBodyParam) (string, error) {
	verifyKey := random.GetStr(100)
	return verifyKey, db.Create(&models.UserVerificationRequest{
		UserID:          user.ID,
		Type:            "EMAIL",
		VerificationKey: verifyKey,
	}).Error
}

func defineReferralUser(db *gorm.DB, code string, user *models.User) error {
	var referralCode models.ReferralCode
	if db.Preload("User").Where(&models.ReferralCode{
		Code: code,
	}).First(&referralCode).RecordNotFound() {
		return fmt.Errorf("Invalid referral code")
	}

	if referralCode.User.BlockedAt != nil {
		return fmt.Errorf("Referrer User Blocked")
	}

	return db.Create(&models.ReferralUser{
		UserID:         user.ID,
		ReferralCodeID: referralCode.ID,
	}).Error
}

func generateMailVerifyTpl(name, token string) (string, string) {
	webBaseURL := os.Getenv("WEB_BASE_URL")
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Ayok.be",
			Link: webBaseURL,
			// Optional product logo
			Logo: os.Getenv("MAIL_LOGO_IMG_URL"),
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Hi " + name + "! let's verify your account.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To verify, please click the link below:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Verify My Account",
						Link:  webBaseURL + "/verify?token=" + token,
					},
				},
			},
			Outros: []string{
				"This email is generated by system, please not replying to this email",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	return emailBody, emailText
}

func sendEmailVerification(name, email, token string) error {
	msgHTML, msgText := generateMailVerifyTpl(name, token)
	out, err := mail.Send(&mail.Template{
		From: os.Getenv("SES_FROM_EMAIL"),
		Recipients: []*string{
			&email,
		},
		Subject: "Verify Your Account",
		HTML:    msgHTML,
		Text:    msgText,
	})

	logrus.Info(out)
	return err
}
