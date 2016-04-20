package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/db"
	"github.com/philippecarle/go-user-api/middlewares"
	"github.com/philippecarle/go-user-api/routing"
	"os"
	"runtime"
)

const (
	// Port at which the server starts listening
	Port = "8888"
)

// Init DB, CPU Numbers, Gin Mode
func init() {
	db.Connect()
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.New()

	setMiddlewares(r)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	routing.New(r)

	endless.ListenAndServe(":"+port, r)
}

func setMiddlewares(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Connect)
	r.Use(middlewares.CORS())
}