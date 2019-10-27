package context

import (
	"github.com/jinzhu/gorm"
	"github.com/fasthttp/session"
)

type AppContext struct {
	DB *gorm.DB
	Sesssion *session.Session
}