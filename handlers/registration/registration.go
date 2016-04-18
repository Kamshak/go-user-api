package registration

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/encryption"
	"github.com/philippecarle/go-user-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"unicode"
)

type RegisterForm struct {
	Username string   `form:"username" binding:"required"`
	Password string   `form:"password" binding:"required"`
	Roles    []string `form:"roles[]"`
}

var mustHave = []func(rune) bool{
	unicode.IsUpper,
	unicode.IsLower,
	unicode.IsPunct,
	unicode.IsDigit,
}

// Register a user
// http -f POST localhost:8888/admin/users "Authorization:Bearer XXXXXXXXXXXX" username=yeah@hi.tld password=myPa$$W0rd roles=ADMIN
func RegisterHandler(c *gin.Context) {

	db := c.MustGet("db").(*mgo.Database)

	var form RegisterForm

	err := c.Bind(&form)

	log.Print(form)

	if err != nil {
		// TODO loop through errors and get missing fields names
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some fields are missing",
		})
	} else {
		formErrors := validateForm(form)

		if len(formErrors) > 0 {
			e := make([]string, len(formErrors)-1)
			for _, err := range formErrors {
				e = append(e, err.Error())
			}
			log.Print(formErrors)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Some values are incorrect",
				"errors":  e,
			})
		} else {
			salt, hash := encryption.EncryptPassword(form.Password)

			u := users.GetUserByUserName(form.Username)

			if u.Username == form.Username {
				c.JSON(http.StatusConflict, gin.H{
					"message": "Username already used",
				})
			} else {

				user := users.User{bson.NewObjectId(), form.Username, salt, hash, form.Roles}

				db.C(users.UsersCollection).Insert(user)

				c.JSON(http.StatusCreated, user)
			}
		}
	}
}

func validateForm(f RegisterForm) []error {

	var e []error

	if !govalidator.IsEmail(f.Username) {
		e = append(e, errors.New("Username must be an email"))
	}

	if !passwordOK(f.Password) {
		e = append(e, errors.New("Password must contains uppercase and lowercase, special characters and digits"))
	}

	return e
}

func passwordOK(p string) bool {
	for _, testRune := range mustHave {
		found := false
		for _, r := range p {
			if testRune(r) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
