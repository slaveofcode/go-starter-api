package context

import (
	"github.com/fasthttp/session"
	"github.com/jinzhu/gorm"
)

// AppContext application context
type AppContext struct {
	DB      *gorm.DB
	Session *session.Session
}
