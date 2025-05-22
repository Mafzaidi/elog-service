package auth

import (
	"github.com/mafzaidi/elog/config"
	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/pkg/authorizer"
)

type UseCase interface {
	Register(pl *RegisterPayload) error
	Login(email, password, validToken string, conf *config.Config) (*UserToken, error)
	ConfirmPassword(userID string, password string) error
}

type UserToken struct {
	User   *entities.User
	Token  string
	Claims *authorizer.Claims
}
