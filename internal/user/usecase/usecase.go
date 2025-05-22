package usecase

import (
	"errors"

	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUC struct {
	repo user.Repository
}

func NewUserUseCase(repo user.Repository) user.UseCase {
	return &UserUC{
		repo: repo,
	}
}

func (u *UserUC) User(ID primitive.ObjectID) (*entities.User, error) {

	user, err := u.repo.FindByID(ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
