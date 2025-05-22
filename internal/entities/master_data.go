package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MasterData struct {
	ID         primitive.ObjectID `json:"id"`
	Group      string             `json:"group"`
	Key        string             `json:"key"`
	Attributes struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"attributes"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	IsActive  bool      `json:"is_active"`
}
