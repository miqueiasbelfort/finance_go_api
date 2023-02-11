package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follow struct {
	UserId   primitive.ObjectID `json:"userId" bson:"userId"`
	FullName string             `json:"fullname" bson:"fullname"`
}

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Following []Follow           `json:"following" bson:"following"`
	Followers []Follow           `json:"followers" bson:"followers"`
}
