package service

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	ServicesHasAccount(userID primitive.ObjectID, isActive *bool) ([]entities.MasterData, error)
}
