package controllers

import (
	"api/src/cryptography"
	"api/src/database"
	"api/src/models"
	"api/src/verifications"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAUser(w http.ResponseWriter, r *http.Request) {

	client, err := database.ConnectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestVerification := verifications.CreateUser(user)
	if requestVerification {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passwordHash, err := cryptography.Hash(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = primitive.NewObjectID().Timestamp()
	user.UpdatedAt = primitive.NilObjectID.Timestamp()
	user.Password = string(passwordHash)

	collection := client.Database("golang").Collection("users")

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
