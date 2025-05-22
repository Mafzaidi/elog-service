package models

type Event struct {
	EventID int64  `json:"event_id" bson:"eventID"`
	Title   string `json:"title" bson:"title"`
}
