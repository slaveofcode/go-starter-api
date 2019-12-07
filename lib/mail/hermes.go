package mail

import (
	"os"

	"github.com/matcornic/hermes"
)

// GetHermes returns new hermes instance
func GetHermes() *hermes.Hermes {
	return &hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: os.Getenv("SITE_NAME"),
			Link: os.Getenv("WEB_BASE_URL"),
			// Optional product logo
			Logo:      os.Getenv("MAIL_LOGO_IMG_URL"),
			Copyright: os.Getenv("MAIL_COPYRIGHT_FOOTER"),
		},
	}
}
