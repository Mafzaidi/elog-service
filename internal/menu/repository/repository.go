package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MenuRepository struct {
	Collection *mongo.Collection
}

func NewMenuRepository(db *mongo.Database) *MenuRepository {
	return &MenuRepository{
		Collection: db.Collection("menus"),
	}
}

func (r *MenuRepository) FindByID(id primitive.ObjectID) (*entities.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := &models.Menu{}
	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	menu := toEntity(m)
	return menu, nil
}

func (r *MenuRepository) FindManyByFilter(filter bson.M) ([]entities.Menu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var es []entities.Menu

	if filter == nil {
		filter = bson.M{}
	}

	csr, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	for csr.Next(ctx) {
		var m models.Menu
		if err := csr.Decode(&m); err != nil {
			return nil, err
		}
		es = toEntities(es, m)
	}

	if err := csr.Err(); err != nil {
		return nil, err
	}
	menus := es
	return menus, err
}

func toEntity(m *models.Menu) *entities.Menu {
	return &entities.Menu{
		ID:       m.ID,
		Title:    m.Title,
		Url:      m.Url,
		Icon:     m.Icon,
		IsActive: m.IsActive,
		Group:    m.Group,
	}
}

func toEntities(es []entities.Menu, m models.Menu) []entities.Menu {
	es = append(es, entities.Menu{
		ID:       m.ID,
		Title:    m.Title,
		Url:      m.Url,
		Icon:     m.Icon,
		IsActive: m.IsActive,
		Group:    m.Group,
	})
	return es
}
