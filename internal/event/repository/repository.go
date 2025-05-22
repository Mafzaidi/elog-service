package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mafzaidi/elog/internal/event"
	"github.com/mafzaidi/elog/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	Collection *mongo.Collection
}

func NewEventRepository(db *mongo.Database) event.Repository {
	return &EventRepository{
		Collection: db.Collection("events"),
	}
}

func (r *EventRepository) FindByID(id int64) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event := &models.Event{}
	filter := bson.M{"eventID": id}
	err := r.Collection.FindOne(ctx, filter).Decode(event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return event, nil
}
