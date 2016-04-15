package users

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/philippecarle/auth/encryption"
	"github.com/philippecarle/auth/models"
	"log"
	"github.com/philippecarle/auth/jwt"
)

// Register a user
func Register (c *gin.Context) {

	db := c.MustGet("db").(*mgo.Database)

	username := c.PostForm("username")
	password := c.PostForm("password")

	salt, hash := encryption.EncryptPassword(password)

	existingUser := models.User{}
	db.C("users").Find(bson.M{"username": username}).One(&existingUser)

	if existingUser.Username == username {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Username already used",
		})
	}

	user := models.User{bson.NewObjectId(), username, salt, hash}

	db.C("users").Insert(user)

	c.JSON(http.StatusCreated, user)
}

// Validate Credentials and return a JWT
func Login (c *gin.Context) {

	var json models.Login
	db := c.MustGet("db").(*mgo.Database)

	err := c.Bind(&json)

	if err == nil {
		user := models.User{}

		db.C("users").Find(bson.M{"username": json.Username}).One(&user)

		verification, err := encryption.IsPasswordValid(json.Password, string(user.Salt), string(user.Hash))

		if err != nil {
			log.Fatal(err)
		}

		if verification {
			c.JSON(http.StatusOK, authjwt.GenerateJWT(user))
		} else {
			c.JSON(http.StatusUnauthorized, []int{})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields username and password are mandatory",
		})
	}
}

func Me (c *gin.Context) {

}