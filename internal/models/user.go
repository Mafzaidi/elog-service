package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Fullname     string             `bson:"fullName"`
	PhoneNumber  string             `bson:"phoneNumber"`
	Password     string             `bson:"password,omitempty"`
	Email        string             `bson:"email"`
	Group        string             `bson:"group"`
	MasterKeyEnc string             `bson:"masterKeyEncrypted"`
	Salt         string             `bson:"salt"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
