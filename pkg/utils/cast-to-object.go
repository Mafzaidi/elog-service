package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectID(hex string) (*primitive.ObjectID, error) {
	if hex == "" {
		return nil, nil
	}
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
