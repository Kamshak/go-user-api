package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/auth/db"
	"github.com/philippecarle/auth/middlewares"
	"github.com/philippecarle/auth/handlers/users"
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
	router := gin.Default()

	router.Use(middlewares.Connect)

	router.POST("/register", users.Register)
	router.POST("/login", users.Login)
	router.GET("/me", users.Me)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	router.Run(":" + port)
}
