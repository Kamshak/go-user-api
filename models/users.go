package users

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/philippecarle/auth/db"
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
}

func GetUserByUserName(username string) (User) {
	s := db.Session.Clone()
	defer s.Close()

	user := User{}
	s.DB(db.Mongo.Database).C(UsersCollection).Find(bson.M{"username": username}).One(&user)

	return user
}