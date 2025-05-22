package menu

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Menu(id primitive.ObjectID) (*entities.Menu, error)
	ActiveMenus(isActive *bool) ([]entities.Menu, error)
	ActiveUserMenus(isActive *bool, group string) ([]entities.Menu, error)
}
