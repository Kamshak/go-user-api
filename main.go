package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/db"
	"github.com/philippecarle/go-user-api/handlers/login"
	"github.com/philippecarle/go-user-api/handlers/registration"
	"github.com/philippecarle/go-user-api/handlers/user"
	"github.com/philippecarle/go-user-api/middlewares"
	"os"
	"time"
)

const (
	// Port at which the server starts listening
	Port = "8888"
)

func init() {
	db.Connect()
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Connect)

	gin.SetMode(gin.DebugMode)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "User API",
		Key:           []byte("54xDEBGEMtnNZJGpPahYzdd47nuQ8M64QpXeDyLnGAH3Gq3HQwnbRG625z9pvNAgSrgp5vTrpC7u2bcqfDs23WX93tefUf8dp7aqxyQVZFzzKhsGtmHgA29r"),
		Timeout:       time.Hour * 72,
		Authenticator: login.LoginHandler,
		Authorizator: func(userId string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}

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

	endless.ListenAndServe(":"+port, r)
}
