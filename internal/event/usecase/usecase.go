package usecase

import (
	"github.com/mafzaidi/elog/internal/event"
	"github.com/mafzaidi/elog/internal/models"
)

type EventUC struct {
	repo event.Repository
}

func NewEventUseCase(repo event.Repository) event.UseCase {
	return &EventUC{
		repo: repo,
	}
}

func (u *EventUC) FindByID(id int64) (*models.Event, error) {
	event, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}
