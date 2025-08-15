package auth

import (
	"github.com/mafzaidi/elog/config"
	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/pkg/authorizer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Register(pl *RegisterPayload) error
	Login(email, password, validToken string, conf *config.Config) (*UserToken, error)
	ConfirmPassword(userID string, password string) error
	User(ID primitive.ObjectID) (*entities.User, error)
}

type UserToken struct {
	User   *entities.User
	Token  string
	Claims *authorizer.Claims
}
