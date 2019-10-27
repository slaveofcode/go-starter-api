package route

import (
	"github.com/fasthttp/router"
	"github.com/slaveofcode/go-starter-api/handlers"
	"github.com/slaveofcode/go-starter-api/handlers/user"
	"github.com/slaveofcode/go-starter-api/context"
)

// New route creation
func New(appCtx *context.AppContext) *router.Router {
	router := router.New()

	router.NotFound = NotFoundHandler
	router.PanicHandler = PanicHandler

	router.GET("/", handlers.Pinger)

	userHandlers := user.NewUser(appCtx)
	router.GET("/users", userHandlers.List)

	return router
}
