package verifications

import "api/src/models"

func CreateUser(user models.User) bool {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return true
	}
	return false
}
