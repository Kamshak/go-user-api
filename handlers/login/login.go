package login

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/philippecarle/auth/models"
	"github.com/philippecarle/auth/encryption"
	"github.com/philippecarle/auth/jwt"
)

// Binding from JSON
type LoginForm struct {
	Username string `form:"_username" binding:"required"`
	Password string `form:"_password" binding:"required"`
}


// Validate Credentials and return a JWT
func Login(c *gin.Context) {

	var form LoginForm

	err := c.Bind(&form)

	if err == nil {
		u := users.GetUserByUserName(form.Username)

		verification, err := encryption.IsPasswordValid(form.Password, string(u.Salt), string(u.Hash))

		if err != nil {
			log.Fatal(err)
		}

		if verification {
			c.JSON(http.StatusOK, jwt.GenerateJWT(u))
		} else {
			c.JSON(http.StatusUnauthorized, []int{})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields username and password are mandatory",
		})
	}
}