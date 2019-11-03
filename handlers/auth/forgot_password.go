package auth

import (
	"encoding/json"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/matcornic/hermes"
	"github.com/sirupsen/logrus"
	"github.com/slaveofcode/go-starter-api/lib/httpresponse"
	"github.com/slaveofcode/go-starter-api/lib/mail"
	"github.com/slaveofcode/go-starter-api/lib/random"
	"github.com/slaveofcode/go-starter-api/repository/pg/models"
	"github.com/valyala/fasthttp"
	validator "gopkg.in/go-playground/validator.v9"
)

type forgotPasswordBodyParam struct {
	Email string `json:"email" validate:"required,email"`
}

// ForgotPassword handles user forgot password request
func (auth Auth) ForgotPassword(ctx *fasthttp.RequestCtx) {
	var param forgotPasswordBodyParam

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

	var user models.User
	db.Model(&cred).Related(&user)

	logrus.Info(user.Name)

	var resetCred models.ResetCredential
	if db.Where(&models.ResetCredential{
		CredentialID: cred.ID,
		IsExpired:    false,
	}).First(&resetCred).RecordNotFound() {
		rec, err := createResetCredential(db, &cred)
		if err != nil {
			httpresponse.JSONErr(ctx, "Unable to process forgot password: "+err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		sendEmail(user.Name, cred.Email, rec.ResetToken)

		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}

	now := time.Now()
	if resetCred.ValidUntil.After(now) {
		httpresponse.JSONOk(ctx, fasthttp.StatusOK)
		return
	}

	if err := db.Model(&resetCred).Updates(&models.ResetCredential{
		IsExpired: true,
	}).Error; err != nil {
		httpresponse.JSONErr(ctx, "Unable to set expiry on existing record", fasthttp.StatusInternalServerError)
		return
	}

	rec, err := createResetCredential(db, &cred)
	if err != nil {
		httpresponse.JSONErr(ctx, "Unable to process forgot password: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	sendEmail(user.Name, cred.Email, rec.ResetToken)

	httpresponse.JSONOk(ctx, fasthttp.StatusOK)
	return
}

func createResetCredential(db *gorm.DB, cred *models.Credential) (*models.ResetCredential, error) {
	token := random.GetStr(128)

	reset := models.ResetCredential{
		CredentialID: cred.ID,
		ResetToken:   token,
		ValidUntil:   time.Now().Add(2 * time.Hour), // 2 hours from now
	}

	if err := db.Create(&reset).Error; err != nil {
		return nil, err
	}

	return &reset, nil
}

func generateMailTpl(name, token string) (string, string) {
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Ayok.be",
			Link: "https://ayok.be/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Hi " + name + "! Let's reset your password.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To reset your password, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Reset My Password",
						Link:  "https://ayok.be/reset?token=" + token,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
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

func sendEmail(name, email, token string) error {
	msgHTML, msgText := generateMailTpl(name, token)
	out, err := mail.Send(&mail.Template{
		From: os.Getenv("SES_FROM_EMAIL"),
		Recipients: []*string{
			&email,
		},
		Subject: "Reset Password Account",
		HTML:    msgHTML,
		Text:    msgText,
	})

	logrus.Info(out)
	return err
}
