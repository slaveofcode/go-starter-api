package route

import (
	"github.com/fasthttp/router"
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/slaveofcode/go-starter-api/handlers"
	"github.com/slaveofcode/go-starter-api/handlers/auth"
	"github.com/slaveofcode/go-starter-api/handlers/user"
)

// New route creation
func New(appCtx *context.AppContext) *router.Router {
	router := router.New()

	router.NotFound = NotFoundHandler
	router.PanicHandler = PanicHandler

	router.GET("/", handlers.Pinger)

	// Auth Handlers
	authHandlers := auth.NewAuth(appCtx)
	router.POST("/auth/register", authHandlers.Register)
	router.POST("/auth/verify", authHandlers.Verify)
	router.POST("/auth/forgot", authHandlers.ForgotPassword)
	router.POST("/auth/reset_check", authHandlers.ResetPasswordCheck)
	router.POST("/auth/reset", authHandlers.ResetPassword)
	router.POST("/auth/login", authHandlers.Login)
	router.POST("/auth/logout", authHandlers.Logout)

	// User Handlers
	userHandlers := user.NewUser(appCtx)
	router.GET("/users", userHandlers.List)

	return router
}
