package account

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePayload struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Notes       string `json:"notes"`
	Service     string `json:"service"`
	PasswordApp string `json:"password_app"`
	IsActive    *bool  `json:"is_active"`
}

type CreateParams struct {
	UserID      primitive.ObjectID
	PasswordApp string
	Username    string
	Password    string
	Host        string
	Notes       string
	Service     string
	IsActive    *bool
}

type FilterUsersAccountsPayload struct {
	IsActive *bool `query:"is_active"`
}

type FilterUsersAccountsResponse struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Service struct {
		ID   string `json:"id"`
		Code string `json:"code"`
		Key  string `json:"key"`
		Name string `json:"name"`
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

type ServiceParams struct {
	Group    string
	Key      string
	IsActive *bool
}
