package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterPayload struct {
	Username    string `json:"username" validate:"required"`
	FullName    string `json:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type (
	LoginPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		ID          primitive.ObjectID `json:"id"`
		Username    string             `json:"username"`
		Fullname    string             `json:"full_name"`
		PhoneNumber string             `json:"phone_number"`
		Email       string             `json:"email"`
		Group       string             `json:"group"`
		CreatedAt   time.Time          `json:"created_at"`
		UpdatedAt   time.Time          `json:"updated_at"`
		AccessToken struct {
			Type      string    `json:"type"`
			Token     string    `json:"token"`
			ExpiresAt time.Time `json:"expires_at"`
		} `json:"access_token"`
	}
)

type GetCurrentUserResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Username    string             `json:"username"`
	Fullname    string             `json:"full_name"`
	PhoneNumber string             `json:"phone_number"`
	Email       string             `json:"email"`
	Group       string             `json:"group"`
}
