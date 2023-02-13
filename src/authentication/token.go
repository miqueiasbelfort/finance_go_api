package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAToken(userID primitive.ObjectID) (string, error) {

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("123456789"))
}

// get the token and verify if the token is valid
func ValidateToken(r *http.Request) error {
	tokenString := getAToken(r)
	token, err := jwt.Parse(tokenString, getAVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token is Invalid")

}

// Get the token by request
func getAToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	fmt.Println("Don't have a token")
	return ""
}

func getAVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method! %v", token.Header["alg"])
	}
	return "123456789", nil
}

func GetUserIDbyToken(r *http.Request) (interface{}, error) {
	tokenString := getAToken(r)
	token, erro := jwt.Parse(tokenString, getAVerificationKey)
	if erro != nil {
		return primitive.NilObjectID, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := permissions["userID"]
		return userID, nil
	}

	return nil, errors.New("Token is invalid")
}
