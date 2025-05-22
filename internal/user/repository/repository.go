package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/models"
	"github.com/mafzaidi/elog/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) user.Repository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (r *UserRepository) FindByID(ID primitive.ObjectID) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := &models.User{}
	filter := bson.M{"_id": ID}

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
