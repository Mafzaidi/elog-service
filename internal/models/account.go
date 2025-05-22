package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  primitive.ObjectID `bson:"userID"`
	Service struct {
		ID   primitive.ObjectID `bson:"_id,omitempty"`
		Code string             `bson:"code"`
		Key  string             `bson:"key"`
		Name string             `bson:"name"`
	} `bson:"service"`
	Username          string    `bson:"username"`
	PasswordEncrypted string    `bson:"passwordEncrypted"`
	Salt              string    `bson:"salt"`
	Host              string    `bson:"host"`
	Notes             string    `bson:"notes"`
	IsActive          bool      `bson:"isActive"`
	CreatedAt         time.Time `bson:"createdAt"`
	UpdatedAt         time.Time `bson:"updatedAt"`
}
