package registration

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/go-user-api/db"
	"github.com/philippecarle/go-user-api/models"
	"github.com/philippecarle/go-user-api/password"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type RegisterForm struct {
	Username string   `form:"username" binding:"required"`
	Password string   `form:"password" binding:"required"`
	Roles    []string `form:"roles[]"`
}

const PASSWORD_REQUIREMENT = "Password must contains uppercase and lowercase, special characters and digits"
const USERNAME_REQUIREMENT = "Username must be an email"

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
			salt, hash := password.EncryptPassword(form.Password)

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

// Change User password
func ChangePasswordHandler(c *gin.Context) {
	s := db.Session.Clone()
	defer s.Close()

	payload, _ := c.Get("JWT_PAYLOAD")
	p, _ := payload.(map[string]interface{})

	user := users.GetUserByUserName(p["id"].(string))

	current := c.Query("current_password")
	new := c.Query("new_password")

	valid, _ := password.IsPasswordValid(current, string(user.Salt), string(user.Hash))

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Current password is incorrect",
		})
	} else {
		if password.CheckPasswordRequirements(new) {
			user.Salt, user.Hash = password.EncryptPassword(new)
			s.DB(db.Mongo.Database).C(users.UsersCollection).Update(bson.M{"username": p["id"].(string)}, user)
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": PASSWORD_REQUIREMENT,
			})
		}
	}
}

// Validate register form
func validateForm(f RegisterForm) []error {

	var e []error

	if !govalidator.IsEmail(f.Username) {
		e = append(e, errors.New(USERNAME_REQUIREMENT))
	}

	if !password.CheckPasswordRequirements(f.Password) {
		e = append(e, errors.New(PASSWORD_REQUIREMENT))
	}

	return e
}
