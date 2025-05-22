package models

type Menu struct {
	ID       string `bson:"_id,omitempty"`
	Title    string `bson:"title"`
	Url      string `bson:"url"`
	Icon     string `bson:"icon"`
	IsActive bool   `bson:"isActive"`
	Group    string `bson:"group"`
}
