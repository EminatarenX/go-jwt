package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct{
	Email string `json: "email"`
	Password string `json: "password"`
}

type User struct{
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	Token string `json:"token"`
	CreatedAt string `json:"created_at"`
}
