package login

import (
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/models"
	"github.com/philippecarle/go-user-api/password"
)

// Validate Credentials and return a JWT
// http --json POST localhost:8888/login username=yeah@hi.tld password=myPa$$W0rd
func LoginHandler(username string, pw string, c *gin.Context) (string, bool) {

	u := users.GetUserByUserName(username)

	verification, _ := password.IsPasswordValid(pw, string(u.Salt), string(u.Hash))

	if verification {
		return username, true
	} else {
		return username, false
	}
}
