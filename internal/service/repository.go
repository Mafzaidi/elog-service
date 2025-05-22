package service

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	FindByFilter(filter bson.M) (*entities.MasterData, error)
	FindManyByFilter(filter bson.M) ([]entities.MasterData, error)
}
