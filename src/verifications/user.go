package verifications

import (
	"api/src/models"
	"errors"
)

func CreateUser(user models.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return errors.New("Fill all fields to create a user")
	}
	return nil
}
func UpdateUser(user models.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return errors.New("Fill all fields to update a user")
	}
	return nil
}
