package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRepository struct {
	Collection *mongo.Collection
}

func NewServiceRespository(db *mongo.Database) *ServiceRepository {
	return &ServiceRepository{
		Collection: db.Collection("master_data"),
	}
}

func (r *ServiceRepository) FindByFilter(filter bson.M) (*entities.MasterData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = bson.M{}
	}

	m := &models.MasterData{}

	err := r.Collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("service not found")
		}
		return nil, err
	}

	service := toEntity(m)

	return service, nil
}

func (r *ServiceRepository) FindManyByFilter(filter bson.M) ([]entities.MasterData, error) {
	return nil, nil
}

func toEntity(m *models.MasterData) *entities.MasterData {
	return &entities.MasterData{
		ID:    m.ID,
		Group: m.Group,
		Key:   m.Key,
		Attributes: struct {
			Code string `json:"code"`
			Name string `json:"name"`
		}(m.Attributes),
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedAt: m.UpdatedAt,
		UpdatedBy: m.UpdatedBy,
		IsActive:  m.IsActive,
	}
}
