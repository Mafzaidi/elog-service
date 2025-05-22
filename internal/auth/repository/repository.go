package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mafzaidi/elog/internal/auth"
	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	Collection *mongo.Collection
}

func NewAuthRepository(db *mongo.Database) auth.Repository {
	return &AuthRepository{
		Collection: db.Collection("users"),
	}
}

func (r *AuthRepository) FindByID(id primitive.ObjectID) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := &models.User{}
	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user := toEntity(m)

	return user, nil
}

func (r *AuthRepository) FindByUsername(username string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := &models.User{}
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&m)

	user := toEntity(m)
	return user, err
}

func (r *AuthRepository) FindByEmail(email string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := &models.User{}
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&m)

	user := toEntity(m)
	return user, err
}

func (r *AuthRepository) Create(user *entities.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := toModel(user)

	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(ctx, m)
	return err
}

func toEntity(m *models.User) *entities.User {
	return &entities.User{
		ID:           m.ID,
		Username:     m.Username,
		Fullname:     m.Fullname,
		PhoneNumber:  m.PhoneNumber,
		Password:     m.Password,
		Email:        m.Email,
		Group:        m.Group,
		MasterKeyEnc: m.MasterKeyEnc,
		Salt:         m.Salt,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func toModel(e *entities.User) *models.User {
	return &models.User{
		Username:     e.Username,
		Fullname:     e.Fullname,
		PhoneNumber:  e.PhoneNumber,
		Password:     e.Password,
		Email:        e.Email,
		Group:        e.Group,
		MasterKeyEnc: e.MasterKeyEnc,
		Salt:         e.Salt,
	}
}
