package menu

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	FindByID(id primitive.ObjectID) (*entities.Menu, error)
	FindManyByFilter(filter bson.M) ([]entities.Menu, error)
}
