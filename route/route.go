package route

import (
	"github.com/fasthttp/router"
	"github.com/jinzhu/gorm"
	"github.com/slaveofcode/go-starter-api/handlers"
)

// New route creation
func New(db *gorm.DB) *router.Router {
	router := router.New()

	router.NotFound = NotFoundHandler
	router.PanicHandler = PanicHandler

	router.GET("/", handlers.Pinger)

	userHandlers := handlers.NewUser(db)
	router.GET("/users", userHandlers.List)

	return router
}
