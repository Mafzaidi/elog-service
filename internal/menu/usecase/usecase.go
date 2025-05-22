package usecase

import (
	"errors"

	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/menu"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuUC struct {
	repo menu.Repository
}

func NewMenuUseCase(repo menu.Repository) menu.UseCase {
	return &MenuUC{
		repo: repo,
	}
}

func (u *MenuUC) Menu(id primitive.ObjectID) (*entities.Menu, error) {

	menu, err := u.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}

	return menu, nil
}

func (u *MenuUC) ActiveMenus(isActive *bool) ([]entities.Menu, error) {
	filter := bson.M{}

	if isActive != nil {
		filter["isActive"] = *isActive
	}

	menus, err := u.repo.FindManyByFilter(filter)
	if err != nil {
		return nil, errors.New("menus not found")
	}

	return menus, nil
}

func (u *MenuUC) ActiveUserMenus(isActive *bool, group string) ([]entities.Menu, error) {
	filter := bson.M{}

	if isActive != nil {
		filter["isActive"] = *isActive
	}

	if group != "" {
		filter["group"] = group
	}

	menus, err := u.repo.FindManyByFilter(filter)
	if err != nil {
		return nil, errors.New("menus not found")
	}

	return menus, nil
}
