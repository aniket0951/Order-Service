package helper

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidatePrimitiveId(userId string) (primitive.ObjectID, error) {

	if !primitive.IsValidObjectID(userId) {
		return primitive.NewObjectID(), errors.New("invalid id received")
	}
	return ConvertStringToPrimitive(userId)
}

func ConvertStringToPrimitive(id string) (primitive.ObjectID, error) {
	objId, objErr := primitive.ObjectIDFromHex(id)

	if objErr != nil {
		return primitive.NewObjectID(), objErr
	}

	return objId, nil
}
