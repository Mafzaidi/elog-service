package account

import (
	"github.com/mafzaidi/elog/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Store(pl *CreateParams) error
	UserAccounts(userID primitive.ObjectID, isActive *bool) ([]entities.Account, error)
}
