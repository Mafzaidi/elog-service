package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetPayload struct {
	ID string `query:"id"`
}

type GetResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Username    string             `json:"username"`
	Fullname    string             `json:"full_name"`
	PhoneNumber string             `json:"phone_number"`
	Email       string             `json:"email"`
	Group       string             `json:"group"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
