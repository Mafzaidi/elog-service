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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountRepository struct {
	Collection *mongo.Collection
}

func NewAccountRepository(db *mongo.Database) *AccountRepository {
	return &AccountRepository{
		Collection: db.Collection("accounts"),
	}
}

func (r *AccountRepository) Create(account *entities.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := toModel(account)
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(ctx, m)

	return err
}

func (r *AccountRepository) Upsert(filter bson.M, account *entities.Account) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = bson.M{}
	}
	m := toModel(account)
	update := bson.M{
		"$set": bson.M{
			"passwordEncrypted": m.PasswordEncrypted,
			"salt":              m.Salt,
			"host":              m.Host,
			"notes":             m.Notes,
			"updatedAt":         time.Now(),
		},
		"$setOnInsert": bson.M{
			"_id":       primitive.NewObjectID(),
			"createdAt": time.Now(),
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := r.Collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *AccountRepository) FindByID(id primitive.ObjectID) (*entities.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := &models.Account{}
	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	account := toEntity(m)

	return account, nil
}

func (r *AccountRepository) FindByFilter(filter bson.M) (*entities.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = bson.M{}
	}

	m := &models.Account{}

	err := r.Collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("service not found")
		}
		return nil, err
	}

	account := toEntity(m)

	return account, nil
}

func toEntity(m *models.Account) *entities.Account {
	return &entities.Account{
		ID:     m.ID,
		UserID: m.UserID,
		Service: struct {
			ID   primitive.ObjectID `json:"id,omitempty"`
			Code string             `json:"code"`
			Key  string             `json:"key"`
			Name string             `json:"name"`
		}(m.Service),
		Username:          m.Username,
		PasswordEncrypted: m.PasswordEncrypted,
		Salt:              m.Salt,
		Host:              m.Host,
		Notes:             m.Notes,
		IsActive:          m.IsActive,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}

func toModel(e *entities.Account) *models.Account {
	return &models.Account{
		UserID: e.UserID,
		Service: struct {
			ID   primitive.ObjectID `bson:"_id,omitempty"`
			Code string             `bson:"code"`
			Key  string             `bson:"key"`
			Name string             `bson:"name"`
		}(e.Service),
		Username:          e.Username,
		PasswordEncrypted: e.PasswordEncrypted,
		Salt:              e.Salt,
		Host:              e.Host,
		Notes:             e.Notes,
		IsActive:          e.IsActive,
	}
}
