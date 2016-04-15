package authjwt

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/philippecarle/auth/models"
	"io/ioutil"
	"log"
	"time"
)

// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// keys are held in global variables
// i havn't seen a memory corruption/info leakage in go yet
// but maybe it's a better idea, just to store the public key in ram?
// and load the signKey on every signing request? depends on  your usage i guess
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateJWT(user models.User) models.JWT {
	// Create the token
	token := jwt.New(jwt.SigningMethodRS256)
	// Set some claims
	token.Claims["username"] = user.Username
	expires := time.Now().Add(time.Hour * 72)
	token.Claims["expires"] = expires.Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		log.Fatal(err)
	}

	tokenObj := models.JWT{tokenString, expires}

	return tokenObj
}
