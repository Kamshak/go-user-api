package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/auth/db"
	"github.com/philippecarle/auth/handlers/registration"
	"github.com/philippecarle/auth/handlers/login"
	"github.com/philippecarle/auth/middlewares"
	"os"
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

	router.POST("/register", registration.Register)
	router.POST("/login", login.Login)
	//router.GET("/me", users.Me)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	router.Run(":" + port)
}
