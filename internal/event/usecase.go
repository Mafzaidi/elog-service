package event

import "github.com/mafzaidi/elog/internal/models"

type UseCase interface {
	FindByID(id int64) (*models.Event, error)
}
