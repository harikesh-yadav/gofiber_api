package helpers

import (
	"github.com/harikesh-yadav/gofiber_api/database"
	"github.com/harikesh-yadav/gofiber_api/models"
)

func CreateUser(user *models.User) error {

	db, _ := database.Connection()

	if err := db.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func Login(credential *models.Credential) (models.User, error) {
	var user models.User
	db, _ := database.Connection()

	user, err := db.LoginUser(credential)
	if err != nil {
		return user, err
	}

	return user, nil
}
