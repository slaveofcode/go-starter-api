package route

import (
	"github.com/fasthttp/router"
	"github.com/slaveofcode/go-starter-api/context"
	"github.com/slaveofcode/go-starter-api/handlers"
	"github.com/slaveofcode/go-starter-api/handlers/auth"
	"github.com/slaveofcode/go-starter-api/handlers/team"
	"github.com/slaveofcode/go-starter-api/handlers/user"
	"github.com/slaveofcode/go-starter-api/middleware"
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
	router.POST("/users/make-referral-code",
		middleware.AuthenticatedUser(appCtx, userHandlers.MakeReferral),
	)

	// Team Handlers
	teamHandlers := team.NewTeam(appCtx)
	router.POST("/teams/create",
		middleware.AuthenticatedUser(appCtx, teamHandlers.CreateTeam))
	router.POST("/teams/invite",
		middleware.AuthenticatedUser(appCtx, teamHandlers.InviteMember))
	router.POST("/teams/join",
		middleware.AuthenticatedUser(appCtx, teamHandlers.JoinTeam))
	router.POST("/teams/change-role",
		middleware.AuthenticatedUser(appCtx, teamHandlers.ChangeMemberRole))
	router.POST("/teams/freeze-member",
		middleware.AuthenticatedUser(appCtx, teamHandlers.FreezeMember))

	return router
}
