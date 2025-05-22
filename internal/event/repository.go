package event

import "github.com/mafzaidi/elog/internal/models"

type Repository interface {
	FindByID(id int64) (*models.Event, error)
}
