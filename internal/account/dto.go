package account

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreatePayload struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Notes       string `json:"notes"`
	Service     string `json:"service"`
	PasswordApp string `json:"password_app"`
}

type CreateParams struct {
	UserID      primitive.ObjectID
	PasswordApp string
	Username    string
	Password    string
	Host        string
	Notes       string
	Service     string
}

type ServiceParams struct {
	Group    string
	Key      string
	IsActive *bool
}
