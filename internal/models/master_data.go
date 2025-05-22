package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MasterData struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Group      string             `bson:"group"`
	Key        string             `bson:"key"`
	Attributes struct {
		Code string `bson:"code"`
		Name string `bson:"name"`
	} `bson:"attributes"`
	CreatedAt time.Time `bson:"createdAt"`
	CreatedBy string    `bson:"createdBy"`
	UpdatedAt time.Time `bson:"updatedAt"`
	UpdatedBy string    `bson:"updatedBy"`
	IsActive  bool      `bson:"isActive"`
}
