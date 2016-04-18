package user

import (
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/models"
	"net/http"
)

func MeHandler(c *gin.Context) {
	payload, _ := c.Get("JWT_PAYLOAD")
	p, _ := payload.(map[string]interface{})

	user := users.GetUserByUserName(p["id"].(string))

	c.JSON(http.StatusOK, user)
}

func ByUsernameHandler(c *gin.Context) {
	id := c.Param("username")

	var user users.User

	if id == "me" {
		payload, _ := c.Get("JWT_PAYLOAD")
		p, _ := payload.(map[string]interface{})
		user = users.GetUserByUserName(p["id"].(string))
	} else {
		user = users.GetUserByUserName(id)
	}

	c.JSON(http.StatusOK, user)
}

func AllUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users.FindAll())
}
