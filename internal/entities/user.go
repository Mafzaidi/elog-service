package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID
	Username     string
	Fullname     string
	PhoneNumber  string
	Password     string
	Email        string
	Group        string
	MasterKeyEnc string
	Salt         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
