package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/db"
	"github.com/philippecarle/go-user-api/handlers/registration"
	"github.com/philippecarle/go-user-api/handlers/login"
	"github.com/philippecarle/go-user-api/middlewares"
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
