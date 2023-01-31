package controllers

import (
	"api/src/database"
	"api/src/models"
	"context"
	"net/http"
)

func CreateAUser(w http.ResponseWriter, r *http.Request) {

	client, err := database.ConnectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	var user []models.User

	collection := client.Database("golang").Collection("users")

}
