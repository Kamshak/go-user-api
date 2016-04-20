// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/db"
	"github.com/philippecarle/go-user-api/handlers/login"
	"time"
	"github.com/itsjamie/gin-cors"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// Returns the JWT Middleware
func JWT() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
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
}

func CORS() *cors.Middleware {
	return cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, PATH, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	})
}
