package route

import (
	"github.com/fasthttp/router"
	"github.com/slaveofcode/go-starter-api/handlers"
)

// New route creation
func New() *router.Router {
	router := router.New()

	router.GET("/", handlers.Pinger)

	return router
}
