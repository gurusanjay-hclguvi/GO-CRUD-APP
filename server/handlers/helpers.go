package handlers

import (
	"errors"
	"net/http"

	"todo-api/middleware"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUserID(r *http.Request) (primitive.ObjectID, error) {
	val := r.Context().Value(middleware.UserIDKey)
	if val == nil {
		return primitive.NilObjectID, errors.New("user not authenticated")
	}

	userIDStr, ok := val.(string)
	if !ok {
		return primitive.NilObjectID, errors.New("invalid user id type")
	}

	return primitive.ObjectIDFromHex(userIDStr)
}
