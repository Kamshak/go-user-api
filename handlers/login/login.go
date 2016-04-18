package login

import (
	"github.com/philippecarle/go-user-api/encryption"
	"github.com/philippecarle/go-user-api/models"
)

// Validate Credentials and return a JWT
// http --json POST localhost:8888/login username=yeah@hi.tld password=myPa$$W0rd
func LoginHandler(username string, password string) (string, bool) {

	u := users.GetUserByUserName(username)

	verification, _ := encryption.IsPasswordValid(password, string(u.Salt), string(u.Hash))

	if verification {
		return username, true
	} else {
		return username, false
	}
}
