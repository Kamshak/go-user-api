package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/handlers/registration"
	"github.com/philippecarle/go-user-api/handlers/user"
	"github.com/philippecarle/go-user-api/middlewares"
)

// Generates App routing
func New(r *gin.Engine) {
	authMiddleware := middlewares.JWT()

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/users")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/me", user.MeHandler)
		auth.PATCH("/me/change-password", registration.ChangePasswordHandler)
	}

	admin := r.Group("/admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		users := admin.Group("/users")
		users.GET("/:username", user.ByUsernameHandler)
		users.POST("", registration.RegisterHandler)
		users.GET("", user.AllUsersHandler)
	}

	token := r.Group("/token")
	token.Use(authMiddleware.MiddlewareFunc())
	{
		token.GET("/refresh", authMiddleware.RefreshHandler)
	}
}
