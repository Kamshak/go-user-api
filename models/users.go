package users

import (
	"github.com/philippecarle/go-user-api/db"
	"gopkg.in/mgo.v2/bson"
)

const (
	UsersCollection = "users"
)

// User
type User struct {
	Id       bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Salt     []byte        `json:"-" bson:"salt"`
	Hash     []byte        `json:"-" bson:"hash"`
	Roles    []string      `json:"roles" bson:"roles"`
}

// Get a single User by Username
func GetUserByUserName(username string) User {
	s := db.Session.Clone()
	defer s.Close()

	user := User{}
	s.DB(db.Mongo.Database).C(UsersCollection).Find(bson.M{"username": username}).One(&user)

	return user
}

// Find all users
func FindAll() []User {
	s := db.Session.Clone()
	defer s.Close()

	var users []User

	s.DB(db.Mongo.Database).C(UsersCollection).Find(nil).All(&users)

	return users
}
