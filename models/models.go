package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// User
type User struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Salt     []byte        `json:"-" bson:"salt"`
	Hash     []byte        `json:"-" bson:"hash"`
}

// JWT
type JWT struct {
	Token   string          `json:"token"`
	Expires time.Time       `json:"expires"`
}

// Binding from JSON
type Login struct {
	Username     string `form:"_username" binding:"required"`
	Password     string `form:"_password" binding:"required"`
}