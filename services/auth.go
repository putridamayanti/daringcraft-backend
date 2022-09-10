package services

import (
	"daringcraft-backend/database"
	"daringcraft-backend/lib"
	"daringcraft-backend/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func GenerateToken(email string) (*string, error) {
	expire := time.Now().Add(24 * time.Hour)

	claims := models.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err == nil {
		return &tokenString, nil
	}

	return nil, errors.New(err.Error())
}

func CheckHeader(r *http.Request) (string, error) {
	header := r.Header["Authorization"]

	if header == nil {
		return "", errors.New("unauthorized")
	}

	split := strings.Split(header[0], " ")
	if len(split) != 2 || strings.ToLower(split[0]) != "bearer" {
		return "", errors.New("unauthorized")
	}

	return split[1], nil
}

func VerifyToken(tokenString string) (string, error) {
	claims := &models.Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		log.Println(err)
		return "", errors.New("Token invalid")
	}

	return claims.Email, nil
}

func SignIn(params models.Login) (*string, error) {
	user := GetUser(bson.M{"email": params.Email}, nil)

	if user == nil {
		return nil, errors.New("User not found")
	}

	if !user.Active {
		return nil, errors.New("Please confirm your email before go through next process")
	}

	pass := lib.ComparePassword(user.Password, []byte(params.Password))

	if !pass {
		return nil, errors.New("Password does not match")
	}

	token, err := GenerateToken(params.Email)

	if err == nil {
		data := user
		data.LastActive = time.Now()

		_, err = database.UpdateOne(UserCollection, bson.M{"email": params.Email}, data)
		if err != nil {
			return nil, err
		}

		return token, nil
	}

	return nil, errors.New(err.Error())
}
