package login

import (
	"github.com/philippecarle/go-user-api/models"
	"github.com/philippecarle/go-user-api/encryption"
)


// Validate Credentials and return a JWT
// http --json POST localhost:8888/login username=yeah@hi.tld password=myPa$$W0rd
func LoginHandler(userName string, password string) (string, bool) {

	u := users.GetUserByUserName(userName)

	verification, _ := encryption.IsPasswordValid(password, string(u.Salt), string(u.Hash))

	if verification {
		return userName, true
	} else {
		return userName, false
	}
}