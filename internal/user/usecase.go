package user

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	User(ID primitive.ObjectID) (*entities.User, error)
}
