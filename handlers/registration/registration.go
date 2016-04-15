package registration

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/auth/encryption"
	"github.com/philippecarle/auth/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// Register a user
func Register(c *gin.Context) {

	db := c.MustGet("db").(*mgo.Database)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if !govalidator.IsEmail(username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username must be an email",
		})
	} else {
		salt, hash := encryption.EncryptPassword(password)

		u := users.GetUserByUserName(username)

		if u.Username == username {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Username already used",
			})
		} else {
			user := users.User{bson.NewObjectId(), username, salt, hash}

			db.C(users.UsersCollection).Insert(user)

			c.JSON(http.StatusCreated, user)
		}
	}

}