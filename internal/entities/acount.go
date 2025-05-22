package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID      primitive.ObjectID `json:"id,omitempty"`
	UserID  primitive.ObjectID `json:"user_id"`
	Service struct {
		ID   primitive.ObjectID `json:"id,omitempty"`
		Code string             `json:"code"`
		Key  string             `json:"key"`
		Name string             `json:"name"`
	} `json:"service"`
	Username          string    `json:"username"`
	PasswordEncrypted string    `json:"password_encrypted"`
	Salt              string    `json:"salt"`
	Host              string    `json:"host"`
	Notes             string    `json:"notes"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
