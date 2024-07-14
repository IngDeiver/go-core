package userDomain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    FirstName string             `json:"firstName,omitempty"`
    LastName  string             `json:"lastName,omitempty" `
    Email     string             `json:"email,omitempty" `
    Password  string             `json:"password,omitempty" `
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
}

func New() User {
	return User{}
}
