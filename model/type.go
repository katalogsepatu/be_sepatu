package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Fullname        string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Email           string             `json:"email,omitempty" bson:"email,omitempty"`
	Password        string             `json:"password,omitempty" bson:"password,omitempty"`
	ConfirmPassword string             `json:"confirmpassword,omitempty" bson:"confirmpassword,omitempty"`
	PhoneNumber     string             `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Salt            string             `bson:"salt,omitempty" json:"salt,omitempty"`
}

type UpdatePassword struct {
	Oldpassword string `json:"oldpassword,omitempty" bson:"oldpassword,omitempty"`
	Newpassword string `json:"newpassword,omitempty" bson:"newpassword,omitempty"`
}

type KatalogSepatu struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Brand    string             `bson:"brand,omitempty" json:"brand,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
	Price    string             `bson:"price,omitempty" json:"price,omitempty"`
	Color    string             `bson:"color,omitempty" json:"color,omitempty"`
	Diskon   string             `bson:"diskon,omitempty" json:"diskon,omitempty"`
	Image    string             `json:"image,omitempty" bson:"image,omitempty"`
}

type FavoriteSepatu struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Brand    string             `bson:"brand,omitempty" json:"brand,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
	Price    string             `bson:"price,omitempty" json:"price,omitempty"`
	Color    string             `bson:"color,omitempty" json:"color,omitempty"`
	Diskon   string             `bson:"diskon,omitempty" json:"diskon,omitempty"`
	Image    string             `json:"image,omitempty" bson:"image,omitempty"`
}

type Credential struct {
	Status  int    `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Role    string `json:"role,omitempty" bson:"role,omitempty"`
}

type Response struct {
	Status  int    `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	Exp   time.Time          `json:"exp"`
	Iat   time.Time          `json:"iat"`
	Nbf   time.Time          `json:"nbf"`
}
