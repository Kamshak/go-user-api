// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/db"
	"github.com/philippecarle/go-user-api/handlers/login"
	"net/http"
	"time"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// LiberalCORS is a very allowing CORS middleware.
func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

// Returns the JWT Middleware
func JWT() *jwt.GinJWTMiddleware {
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

	return authMiddleware
}
