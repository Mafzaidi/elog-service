package auth

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	FindByID(userID primitive.ObjectID) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Create(user *entities.User) error
}
