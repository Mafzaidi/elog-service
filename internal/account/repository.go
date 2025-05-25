package account

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(account *entities.Account) error
	Upsert(filter bson.M, account *entities.Account) error
	FindByID(id primitive.ObjectID) (*entities.Account, error)
	FindByFilter(filter bson.M) (*entities.Account, error)
	FindManyByFilter(filter bson.M) ([]entities.Account, error)
}
