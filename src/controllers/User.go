package controllers

import (
	"api/src/cryptography"
	"api/src/database"
	"api/src/models"
	"api/src/verifications"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Current   float64            `json:"current" bson:"current`
}

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

func GetAUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var user models.User

	// Data base connection
	client, err := database.ConnectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("golang").Collection("users")

	err = collection.FindOne(context.Background(), models.User{ID: Id}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Current:   user.Current,
	})

}
