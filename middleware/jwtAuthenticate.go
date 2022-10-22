package middleware

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/harikesh-yadav/gofiber_api/models"
)

func CreateJWTToken(user *models.User) (string, error) {

	secret := os.Getenv("JWT_SECRET_KEY")
	minute, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}
	claims["authorization"] = true
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["mobile"] = user.Mobile
	claims["expireAt"] = time.Now().Add(time.Minute * time.Duration(minute)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
