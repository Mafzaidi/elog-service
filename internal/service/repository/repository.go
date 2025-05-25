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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var es []entities.MasterData
	if filter == nil {
		filter = bson.M{}
	}

	csr, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	for csr.Next(ctx) {
		var m models.MasterData
		if err := csr.Decode(&m); err != nil {
			return nil, err
		}
		es = toEntities(es, m)
	}

	if err := csr.Err(); err != nil {
		return nil, err
	}
	services := es
	return services, nil
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

func toEntities(es []entities.MasterData, m models.MasterData) []entities.MasterData {
	es = append(es, entities.MasterData{
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
	})
	return es
}
