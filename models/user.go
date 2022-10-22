package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Age        int       `json:"age" db:"age"`
	Email      string    `json:"email" db:"email"`
	Passcode   string    `json:"passcode" db:"passcode"`
	Address    string    `json:"address" db:"address"`
	Mobile     string    `json:"mobile" db:"mobile"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}

type Credential struct {
	Email    string `json:"email"`
	Passcode string `json:"passcode"`
}
