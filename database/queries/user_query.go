package queries

import (
	"fmt"

	"github.com/harikesh-yadav/gofiber_api/models"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (db *UserQueries) GetUsers() (models.User, error) {
	users := models.User{}
	query := `SELECT * FROM company`

	err := db.Get(&users, query)

	if err != nil {
		return users, err
	}

	return users, nil
}

func (db *UserQueries) CreateUser(user *models.User) error {

	query := "INSERT INTO users (id, name, age , email, passcode, address, mobile, created_at ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := db.Exec(query, user.ID, user.Name, user.Age, user.Email, user.Passcode, user.Address, user.Mobile, user.Created_at)

	if err != nil {
		return err
	}
	return nil
}

func (db *UserQueries) LoginUser(credential *models.Credential) (models.User, error) {
	var user = models.User{}
	var err error

	query := fmt.Sprintf("SELECT * FROM users WHERE email='%s' AND passcode='%s'", credential.Email, credential.Passcode)

	err = db.Get(&user, query)
	if err != nil {
		return user, err
	}

	return user, nil
}
